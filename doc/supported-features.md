## Supported features

<table>
    <tr>
        <td>Cloud-init config field</td>
        <td>Butane config field</td>
        <td>Translation</td>
    </tr>
    <tr>
        <td>groups</td>
        <td>passwd.groups</td>
        <td>cloud-config accepts groups field as a list of strings while butane config needs it as a list of elements each with its
        own name key that stores the group name. This conversion is handled by the transpiler. </td>
    </tr>
    <tr>
        <td>user.name</td>
        <td>passwd.user.name</td>
        <td>direct</td>
    </tr>
    <tr>
        <td>user.gecos</td>
        <td>passwd.user.gecos</td>
        <td>direct</td>
    </tr>
    <tr>
        <td>user.groups</td>
        <td>passwd.user.groups</td>
        <td>cloud-config accepts it as a single string with commas separating each group names while butane config needs it as list in YAML.
        This conversion is handled by the transpiler.
        </td>
    </tr>
    <tr>
        <td>user.shell</td>
        <td>passwd.user.shell</td>
        <td>direct</td>    
    </tr>
    <tr>
        <td>user.uid</td>
        <td>passwd.user.uid</td>
        <td>direct</td>
    </tr>   
    <tr>
        <td>runcmd</td>
        <td>systemd.units</td>
        <td>As per the documentation, butane config does not have a runcmd equivalent. To solve this, The transpiler uses a systemd unit for each
        command. The generated units are enabled and associated with multi-user.target. Each unit uses type as oneshot and the /bin/sh interpreter.</td>    
</tr>
</table>
