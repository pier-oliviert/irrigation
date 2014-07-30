class ZoneBuilder < ActionView::Helpers::FormBuilder
  GPIOS = %w(2 3 4 7 8 9 10 11 14 15 17 18 22 23 24 25 27).map(&:to_i)

  def gpio
    select :gpio do
      available_gpios.map do |gpio|
        selected = gpio == object.gpio
        @template.content_tag :option, gpio, value: gpio, selected: selected
      end.join.html_safe
    end
  end

  private

  def available_gpios
    @available_gpios ||= begin
      pins = GPIOS - Zone.all.pluck(:gpio)
      unless object.gpio.nil?
        pins += [object.gpio]
      end
      pins.sort
    end
  end
end
