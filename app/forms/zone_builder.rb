class ZoneBuilder < ActionView::Helpers::FormBuilder
  GPIOS = %w(0 1 2 3 4 7 8 9 10 11 14 15 17 18 21 22 23 24 25 27).map(&:to_i)

  def gpio
    select :gpio do
      available_gpios.map do |gpio|
        @template.content_tag :option, gpio, value: gpio
      end.join.html_safe
    end
  end

  private

  def available_gpios
    @available_gpios ||= begin
      GPIOS - Zone.all.pluck(:gpio)
    end
  end
end
