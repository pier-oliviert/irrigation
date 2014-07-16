# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20140715133119) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "schedules", force: true do |t|
    t.boolean  "active",     default: true
    t.integer  "day",        default: 0
    t.time     "time",                      null: false
    t.integer  "duration",                  null: false
    t.integer  "zone_id",                   null: false
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  create_table "sprinkles", force: true do |t|
    t.datetime "ends_at"
    t.integer  "zone_id",    null: false
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  create_table "zones", force: true do |t|
    t.integer  "gpio",                       null: false
    t.string   "name",                       null: false
    t.boolean  "opened",     default: false
    t.datetime "created_at"
    t.datetime "updated_at"
  end

end
