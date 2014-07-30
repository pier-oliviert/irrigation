class ZonesController < ApplicationController
  def index
    @zones = Zone.all
    render 'onboarding' and return if @zones.count == 0
  end

  def create
    @zone = Zone.create(zone_params)
    redirect_to :root
  end

  def edit
    @zone = Zone.find(params[:id])
  end

  def update
    @zone = Zone.find(params[:id])
    @zone.update zone_params
    redirect_to :zones
  end

  protected

  def zone_params
    params.require(:zone).permit(:name, :gpio)
  end
end
