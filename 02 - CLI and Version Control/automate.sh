#! /usr/bin/bash
ROOTDIR="$1 at $(date)"
mkdir -p "$ROOTDIR"/{about_me/{personal,professional},my_friends,my_system_info}
echo "https://www.twitter.com/$2" >> "$ROOTDIR/about_me/personal/twitter.txt"
echo "https://www.linkedin.com/in/$3" >> "$ROOTDIR/about_me/professional/linkedin.txt"
uname -a >> "$ROOTDIR/my_system_info/about_this_laptop.txt"
ping -c 3 "google.com" >> "$ROOTDIR/my_system_info/internet_connection.txt"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$ROOTDIR/my_friends/list_of_my_friends.txt"