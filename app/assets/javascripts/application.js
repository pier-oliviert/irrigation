// This is a manifest file that'll be compiled into application.js, which will include all the files
// listed below.
//
// Any JavaScript/Coffee file within this directory, lib/assets/javascripts, vendor/assets/javascripts,
// or vendor/assets/javascripts of plugins, if any, can be referenced here using a relative path.
//
// It's not advisable to add code directly here, but if you do, it'll appear at the bottom of the
// compiled file.
//
// Read Sprockets README (https://github.com/sstephenson/sprockets#sprockets-directives) for details
// about supported directives.
//

//= require jquery
//= require xhr
//= require_tree ./shiny.js
//= require_tree ./zones
//= require_self
//= require listener

$(document).ready(function() {

  $(document).on("click", "a[disabled=disabled]", function(e) {
    e.preventDefault()
    return false
  })

  $(document).on('click', 'a[data-remote=true]:not([disabled])',function(e) {
    if (this == e.target) {
      xhr = new XHR(this)
      xhr.send(this.getAttribute('href'))
      e.preventDefault()
      return false
    }
  })

  $(document).on('submit', '[data-remote=true]:not([disabled])',function(e) {
    if (this == e.target) {
      xhr = new XHR(this)
      data = new FormData(this)
      data.append($('meta[name=csrf-param]').prop('content'), $('meta[name=csrf-token]').prop('content'))
      xhr.send(this.getAttribute('action'), this.getAttribute('method'), data)
      e.preventDefault()
      return false
    }
  })
});
