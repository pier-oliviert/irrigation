class CreateSchedules < ActiveRecord::Migration
  def change
    create_table :schedules do |t|
      t.boolean :active, default: true
      t.integer :day, default: 0
      t.time :time, null: false
      t.integer :duration, null: false
      t.references :zone, null: false
      t.timestamps
    end
  end
end
