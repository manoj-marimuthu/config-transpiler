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
    <tr>
        <td>write_files[].path</td> 
        <td>storage.files[].path</td>
        <td>direct</td>   
    </tr>
    <tr>
        <td>write_files[].content</td>
        <td>storage.files[].contents.inline</td>
        <td>direct</td>
    </tr>
    <tr>
        <td>write_files[].permissions</td>
        <td>storage.files[].mode</td>
        <td>Conversion from octal permission string (such as "0644") to decimal number, handled by the transpiler</td>
    </tr>
    <tr>
        <td>write_files[].owner</td>
        <td>storage.files[].user.name , storage.files[].group.name</td>
        <td>in cloud-config, the write_files[].owner field is given as a string of the form username:groupname 
        whereas in butane config, it uses 2 separate fields to represent username and groupname as listed. The transpiler
        handles this conversion by splitting the string and writing the corresponding fields in butane config structs.</td>
    </tr>
</table>
