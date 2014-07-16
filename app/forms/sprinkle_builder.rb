class SprinkleBuilder < ActionView::Helpers::FormBuilder
  def duration
    select :duration, @template.options_for_select([30, 120, 300])
  end
end
