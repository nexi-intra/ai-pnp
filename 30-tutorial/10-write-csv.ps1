<#---
title: Write CSV
output: test.csv
---
#>

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions") {
    $path = join-path $PSScriptRoot ".." ".."
}
else {
    $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path
$workDir = Join-Path $koksmatDir "workdir"

if (-not (Test-Path $workDir)) {
    New-Item -Path $workDir -ItemType Directory
}




$csv = @"
A,B,C
1,2,3
4,5,6
"@

$array = @()
$array += @{
    A = 1
    B = 2
    C = 3

}
$array += @{
    A = 4
    B = 5
    C = 6

}

# convertto-csv $array -NoTypeInformation | Out-File -FilePath (Join-Path $workDir test2.csv) -Encoding utf8
convertto-json $array  | Out-File -FilePath (Join-Path $workDir test2.json) -Encoding utf8

# $csv | Out-File -FilePath (Join-Path $workDir test.csv) -Encoding utf8