class @XHR
  constructor: (el) ->
    @element(el)
    @element().setAttribute('disabled', true)
    @request = new XMLHttpRequest()
    @request.addEventListener('load', @completed)


  element: (el) ->
    @element = ->
      el

  completed: (e) =>
    @element().removeAttribute('disabled')
    if e.target.responseText.length > 1
      eval(e.target.responseText)(@element())

  send: (src, method = 'GET', data) =>
    @request.open(method, src)
    @request.setRequestHeader('accept', "*/*;q=0.5, #{$.ajaxSettings.accepts.script}")
    @request.setRequestHeader('X-Requested-With', "XMLHttpRequest")

    @request.send(data)
