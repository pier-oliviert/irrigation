class Form
  addEventListeners: =>
    @element().addEventListener "change", @submit
  removeEventListeners: =>
    @element().removeEventListener "change", @submit

  submit: (e) =>
    fd = new FormData(e.target.form)
    fd.append(
      document.querySelector("[name=csrf-param]").getAttribute('content'),
      document.querySelector("[name=csrf-token]").getAttribute('content')
      )
    xhr = new XHR(e.target)
    xhr.send(e.target.form.action, "POST", fd)

    @element().children[0].setAttribute('selected', true)


Shiny.Models.add Form, "Sprinkles.Form"
