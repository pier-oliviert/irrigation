class God
  update: (el) =>
    model = el.getAttribute(Shiny.attributeName)
    if model?
      @create(el, model)
    else
      @destroy(el)

  create: (el) =>
    model = el.getAttribute(Shiny.attributeName)
    if @modelExists(model)
      el.instance = new Shiny.Models.klass[model](el)
      el.instance.element = ->
        el
      el.instance.addEventListeners()
    else
      throw "error: #{model} is not registered. Add your model with Shiny.Models.add(#{model})"

  destroy: (el) =>
    el.instance.removeEventListeners()

  modelExists: (name) =>
    Shiny.Models.klass[name]?
    

Shiny.God = new God
