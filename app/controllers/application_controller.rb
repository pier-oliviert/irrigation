class ApplicationController < ActionController::Base
  # Prevent CSRF attacks by raising an exception.
  # For APIs, you may want to use :null_session instead.
  protect_from_forgery with: :exception

  helper_method :current_user, :signed_in?

  def signed_in?
    !current_user.nil?
  end

  def current_user
    warden.user
  end

  protected

  def osmosis
    env['osmosis']
  end

  def warden
    env['warden']
  end
end
