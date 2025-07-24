# To install run:
# . { iwr -useb https://dl.stackstate.com/stackstate-cli/install.ps1 } | iex; install -StsUrl "url" -StsApiToken "token"
# To uninstall run:
# . { iwr -useb https://dl.stackstate.com/stackstate-cli/install.ps1 } | iex; uninstall
new-module -name "StsCliInstaller" -scriptblock {

  function Error {
    param(
      [String] $msg
    )
    Write-Host "[ERROR]" -ForegroundColor Red -NoNewLine
    Write-Host " $msg"
  }

  function Install-Cli {
    param (
        [string]$StsUrl, # url of the StackState instance to configure (empty means don't configure)
        [string]$StsApiToken, # API-TOKEN of the StackState instance to configure (empty means don't configure)
        [string]$StsCliVersion, # version of the CLI to install (empty means latest)
        [string]$StsCaCertPath, # Path to CA certificate file for HTTPS verification
        [string]$StsCaCertBase64Data, # Base64 encoded CA certificate data for HTTPS verification
        [string]$StsSkipSsl # Skip SSL verification (if set, CA cert options are ignored)
    )
    # Stop on first error
    $ErrorActionPreference = "Stop"

    # Set the source version and target path.
    $CliPath = $env:USERPROFILE +"\stackstate-cli"
    If (!(test-path $CliPath)) {
      md $CliPath > $null
    }
    Write-Host $StsCliVersion
    if (!$StsCliVersion) {
      Invoke-WebRequest https://dl.stackstate.com/stackstate-cli/LATEST_VERSION -OutFile $CliPath\VERSION
      $StsCliVersion=type $CliPath\VERSION
      # The LATEST_VERSION file contains the published tag name, strip the v prefix
      $StsCliVersion=$StsCliVersion.TrimStart("v")
      rm $CliPath\VERSION
    }
    $CliDl = "https://dl.stackstate.com/stackstate-cli/v$StsCliVersion/stackstate-cli-$StsCliVersion.windows-x86_64.zip"

    # Download and unpack the CLI to the target CLI path. Remove remaining artifacts.
    Write-Host "Downloading $CliDl"
    Invoke-WebRequest $CliDl -OutFile $CliPath\stackstate-cli.zip
    Expand-Archive -Path "$CliPath\stackstate-cli.zip" -DestinationPath $CliPath -Force
    rm $CliPath\stackstate-cli.zip

    # Register the CLI path to the current user's PATH, if that wasn't already the case
    $Path = (Get-ItemProperty -Path "Registry::HKEY_CURRENT_USER\Environment" -Name PATH).Path
    if ( $Path -notlike "*$CliPath*" ) {
      $Path = "$Path;$CliPath"
      (Set-ItemProperty -Path "Registry::HKEY_CURRENT_USER\Environment" -Name PATH -Value $Path)
      $MACHINE_PATH = (Get-ItemProperty -Path "Registry::HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\Session Manager\Environment" -Name PATH).path
      # set the current environment path to the new path
      $env:Path = "$Path;$MACHINE_PATH"
    }

    # Verify that 'sts' works
    try {
      & sts version >$null 2>&1
      if ($LastExitCode -ne 0) {
        Error "StackState CLI error code $LastExitCode."
        return
      }
    } Catch {
      Error "Could not find 'sts' on the path."
      return
    }

    # Configure the CLI if config parameters have been set
    if ($StsUrl -and $StsApiToken) {
      if ($StsSkipSsl -eq "true") {
        & sts context save --url $StsUrl --api-token $StsApiToken --skip-ssl $StsSkipSsl
      } elseif ($StsCaCertPath) {
        & sts context save --url $StsUrl --api-token $StsApiToken --ca-cert-path $StsCaCertPath
      } elseif ($StsCaCertBase64Data) {
        & sts context save --url $StsUrl --api-token $StsApiToken --ca-cert-base64-data $StsCaCertBase64Data
      } else {
        & sts context save --url $StsUrl --api-token $StsApiToken
      }
      if ($LastExitCode -ne 0) {
        return
      }
    }

    Write-Host "Success! Type '" -NoNewLine
    Write-Host "sts" -ForegroundColor Green -NoNewline;
    Write-Host "' to get started!"
  }

  function Uninstall-Cli {
    # remove binary
    $CliPath = $env:USERPROFILE+"\stackstate-cli"
    & rm -R $CliPath 2>1  > $null

    # Remove config
    & rm -R $env:USERPROFILE+"\.config\stackstate-cli" 2>1  > $null

    # Remove the CLI from the environment path
    $Path = (Get-ItemProperty -Path "Registry::HKEY_CURRENT_USER\Environment" -Name PATH).Path
    $i = $Path.IndexOf(";$CliPath")
    if ($i -ne -1) {
      $Path = $Path.Remove($i, $CliPath.Length+1)
      (Set-ItemProperty -Path "Registry::HKEY_CURRENT_USER\Environment" -Name PATH -Value $Path)
    }
    Write-Host "Uninstalled the StackState CLI. See you later!"
  }

  set-alias install -value Install-Cli
  set-alias uninstall -value Uninstall-Cli
  export-modulemember -function "Install-Cli" -alias "install"
  export-modulemember -function "Uninstall-Cli" -alias "uninstall"
}
