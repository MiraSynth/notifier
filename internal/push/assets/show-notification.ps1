function Show-Notification {
	[cmdletbinding()]
	Param (
		[string]$context,
		[string]$title,
		[string]$text,
		[string]$icon
	)
	
	[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null;

$xml = @"
<toast>
    <visual>
        <binding template="ToastImageAndText02">
            <image id="1" src="$($icon)" alt="image1"/>
            <text id="1">$($title)</text>
            <text id="2">$($text)</text>
        </binding>  
    </visual>
</toast>
"@
$xmlDocument = [Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime]::New()
$xmlDocument.loadXml($xml)

	$toast = [Windows.UI.Notifications.ToastNotification]::new($xmlDocument);
	$toast.Tag = "PowerShell";
	$toast.Group = "PowerShell";
	$toast.ExpirationTime = [DateTimeOffset]::Now.AddMinutes(1);
	
	$notifier = [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($context);
	$notifier.Show($toast);
}