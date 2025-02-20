# Check if gin is installed
$gin = Get-Command gin -ErrorAction SilentlyContinue
if (-not $gin) {
    Write-Host "❌ 'gin' is not installed. Installing now..."
    go install github.com/codegangsta/gin@latest
}

# Find the Git root (if in a Git project) or use the current directory
if (Test-Path .git) {
    $PROJECT_ROOT = (git rev-parse --show-toplevel)
} else {
    $PROJECT_ROOT = Get-Location
}

Write-Host "Switching to project directory: $PROJECT_ROOT"
Set-Location $PROJECT_ROOT

# Set environment variables
$env:GIN_MODE = "debug"
$env:PORT = "8080"

# Run gin in the project directory
Write-Host "🚀 Starting the Go Gin server with hot reloading..."
cd "./cmd/server"
gin run
