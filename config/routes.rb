Rails.application.routes.draw do
  root 'zones#index'

  resources :zones do
    resources :sprinkles, shallow: true
  end

  resources :schedules
end
