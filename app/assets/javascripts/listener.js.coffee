class Listener
  constructor: ->
    @socket = new WebSocket("ws://#{location.hostname}:21343")
    @socket.onmessage = @received

  addEventListeners: =>

  removeEventListeners: =>

  received: (e) =>
    cmd = JSON.parse(e.data)
    if cmd.status == "open"
      @opened(cmd)
    else
      @closed(cmd)

  opened: (cmd) =>
    z = document.querySelector("#zone_#{cmd.id} div.status")
    z.setAttribute('datetime', cmd.close_at)

  closed: (cmd) =>
    z = document.querySelector("#zone_#{cmd.id} div.status")
    z.removeAttribute('datetime')


Shiny.Models.add Listener, "Listener"
