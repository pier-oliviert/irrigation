class Status
  addEventListeners: =>
    @element().querySelector('select').addEventListener "change", @submit

  removeEventListeners: =>
    @element().querySelector('select').removeEventListener "change", @submit

  submit: (e) =>
    fd = new FormData(e.target.form)
    fd.append(
      document.querySelector("[name=csrf-param]").getAttribute('content'),
      document.querySelector("[name=csrf-token]").getAttribute('content')
      )
    xhr = new XHR(e.target)
    xhr.send(e.target.form.action, "POST", fd)
    e.target.children[0].setAttribute('selected', true)

Shiny.Models.add Status, "Sprinkles.Status"
