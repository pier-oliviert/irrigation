class ZoneDecorator < Cubisme::Decorator::Base
  def list(options)
    content_tag_for :li, record, options do
      [
        content_tag(:h2, link_to(record.name, edit_zone_path(record))),
        status
      ].join.html_safe
    end
  end

  def status
    options = {
      class: %w(status),
      as: 'Sprinkles.Status'
    }

    time = record.closing_at
    unless time.nil?
      options[:datetime] = time
    end

    content_tag :div, options do
      render 'sprinkles/form', zone: record
    end
  end
end
