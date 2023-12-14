param (
	[int]$count = 1000
)

$filename="../benchmark/normalization__large"

$exist = Test-Path $filename
if ($exist) {
    Remove-Item $filename
}

foreach ($i in (1..$count)) {
	Write-Output ("k{0}=v{0}" -f $i) >> $filename
}