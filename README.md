# Picture Frame

## Summary

A picture frame made with a RasPad, Raspberry Pi and Go.

## Details

A Raspberry Pi 3 B+ (really any arm7l pi will work) using a RasPad as a touchscreen. The picture viewer will be in written in Go with Fyne as the GUI. The image syncing will be done with rclone communicating with a folder in NextCloud. A user will save the images to NextCloud and then they will run the syncing application on the picture frame to collect the images. After that process is complete, they will start the picture viewing app. When they want to sync they will close the app with the hidden button on the bottom and run the sync app.

## RasPad

Since this will be used as a frame we are standing the RasPad up instead of the default laying on the desk position. In order for this to work with the Pi 3 B+ we had to create a script to rotate the view and the touchscreen input. The script code is below.
```
#!/usr/bin/bash

xrandr --output 'HDMI-1' --rotate inverted
xinput --set-prop 'ILITEK ILITEK-TP' 'Coordinate Transformation Matrix' -1 0 1 0 -1 1 0 0 1
```
In order to make this auto start you will need a .desktop file in the `/etc/xdg/autostart` diretory.
```
[Desktop Entry]
Version=1.0
Type=Application
Encoding=UTF-8
Name=Rotate Display
Comment=Rotate the display and input
Exec=/home/pi/scripts/rotate.sh
Terminal=false
Categories=Accessories
```

## Desktop Shortcuts

The desktop shortcuts will be available for a user to single tap to open a desktop shortcut. To set the preference for single tap, set it in the filemanager preferences.

* shutdown.desktop
```
[Desktop Entry]
Version=1.0
Type=Application
Encoding=UTF-8
Name=Shutdown
Comment=Shutdown the Display
Exec=sudo /usr/sbin/shutdown -P now
Terminal=false
Categories=Accessories
```

* sync.desktop
```
[Desktop Entry]
Version=1.0
Type=Application
Encoding=UTF-8
Name=Sync
Comment=Sync photos from cloud
Exec=/usr/bin/rclone -P sync remote_name:remote_folder /local/dir
Terminal=true
Categories=Accessories
```
