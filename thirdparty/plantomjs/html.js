system = require('system')
if (system.args.length === 1) {
	console.log('Try to pass some args when invoking this script!');
	phantom.exit();
}
   
url = system.args[1];//获得命令行第二个参数 接下来会用到   
var page = require('webpage').create();   
page.open(url, function (status) {   
    //Page is loaded!   
    if (status !== 'success') {   
        console.log('Unable to post!');   
		phantom.exit();
    } else {   
		window.setTimeout(function () {
			console.log(page.content); 	
			phantom.exit();
        }, 1000); // Change timeout as required to allow sufficient time 
    }      	
});    