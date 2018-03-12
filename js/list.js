var onInit = function() {
    FB.getLoginStatus(function(response) {
        if (response.status === 'connected') {
            getPages()
        }
        else {
            FB.login(getPages, {scope: 'manage_pages'});
        }
    });
}

var app;
$(document).ready(function() {
    app = new Vue({
        el: '#app',
        data: {
            pages: []
        }
    })
})

function getPages() {
    FB.api('/me/accounts', function(response) {
        app.pages = response.data
    })
}