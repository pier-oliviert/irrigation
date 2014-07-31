class SprinkleBuilder < ActionView::Helpers::FormBuilder
  def duration
    select :duration, @template.options_for_select([
      ["Ouverture manuelle", {disabled: true}, -1],
      ["30 Secondes", 30],
      ["2 minutes", 120],
      ["5 minutes", 300]
      ], -1), {}, as: "Sprinkles.Form"
  end
end
