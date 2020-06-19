# duree

_**duree** is my personal very minimalistic and clean browser startpage. 
See it in action below:_


![](demo/demo.gif)


## Installation

- Download the bin/duree binary from this repository and start it with your personal parameters. See usage section.
- In your browser create a bookmark-folder called **duree**.
- Now create new bookmark folders in it. This are your startpage groups.
- Under this groups, put your browser bookmarks.
- That's it :)


## Usage

```
$ ./duree

Usage of ./duree:

  -bookmarkFile string
    	path to the bookmarks.json
  -listenAddr string
    	host:port
```

### Example:
```
$ ./duree --bookmarkFile "/Users/USERNAME/Library/Application Support/Google/Chrome/Default/Bookmarks" --listenAddr "0.0.0.0:3000"
```

### Autostart:

An easy way is to put duree in your crontab with the **@reboot** tag:

```
@reboot /usr/local/bin/duree --bookmarkFile "/Users/USERNAME/Library/Application Support/Google/Chrome/Default/Bookmarks" --listenAddr "0.0.0.0:3000"
```

To open "localhost:3000" on new browser tabs, i personaly use the Chrome extension [New Tab Redirect](https://chrome.google.com/webstore/detail/new-tab-redirect/icpgjfneehieebagbmdbhnlpiopdcmna?hl=de)






