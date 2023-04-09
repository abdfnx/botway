# get latest release
$release_url = "https://api.github.com/repos/abdfnx/botway/releases"
$tag = (Invoke-WebRequest -Uri $release_url -UseBasicParsing | ConvertFrom-Json)[0].tag_name
$loc = "$HOME\AppData\Local\botway"
$url = ""
$arch = $env:PROCESSOR_ARCHITECTURE
$releases_api_url = "https://github.com/abdfnx/botway/releases/download/$tag/botway_windows_${tag}"

if ($arch -eq "AMD64") {
    $url = "${releases_api_url}_amd64.zip"
} elseif ($arch -eq "x86") {
    $url = "${releases_api_url}_386.zip"
} elseif ($arch -eq "arm") {
    $url = "${releases_api_url}_arm.zip"
} elseif ($arch -eq "arm64") {
    $url = "${releases_api_url}_arm64.zip"
}

if (Test-Path -path $loc) {
    Remove-Item $loc -Recurse -Force
}

Write-Host "Installing botway version $tag" -ForegroundColor DarkCyan

Invoke-WebRequest $url -outfile botway_windows.zip

Expand-Archive botway_windows.zip

New-Item -ItemType "directory" -Path $loc

Move-Item -Path botway_windows\bin -Destination $loc

Remove-Item botway_windows* -Recurse -Force

[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$loc\bin", [System.EnvironmentVariableTarget]::User)

if (Test-Path -path $loc) {
    Write-Host "Thanks for installing Botway! Now Refresh your powershell" -ForegroundColor DarkGreen
    Write-Host "If this is your first time using the CLI, be sure to run 'botway --help' first." -ForegroundColor DarkGreen
    Write-Host "Stuck? Join our Discord https://dub.sh/bw-discord" -ForegroundColor DarkCyan
} else {
    Write-Host "Download failed" -ForegroundColor Red
    Write-Host "Please try again later" -ForegroundColor Red
}
