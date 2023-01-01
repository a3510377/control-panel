$cert = New-SelfSignedCertificate -DnsName "a3510377.github.io" -CertStoreLocation cert:\LocalMachine\My -type CodeSigning -Subject "CN=monkey-cat,OU=a3510377.github.io" -NotAfter (Get-Date).AddYears(5)
$pwd = ConvertTo-SecureString -String $Env:password -Force -AsPlainText
Export-PfxCertificate -cert $cert -FilePath mycert.pfx -Password $pwd
