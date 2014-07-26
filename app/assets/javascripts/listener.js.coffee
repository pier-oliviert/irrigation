class Listener
  constructor: ->
    @socket = new WebSocket("ws://#{location.hostname}:21343")
    @socket.onmessage = @received

  addEventListeners: =>

  removeEventListeners: =>


  received: (e) ->
    console.log e

Shiny.Models.add Listener, "Listener"
