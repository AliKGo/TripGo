-- migrations/0001_init.sql
BEGIN;

CREATE TABLE roles (value text PRIMARY KEY);
INSERT INTO roles (value) VALUES ('PASSENGER'), ('DRIVER'), ('ADMIN');

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  email varchar(100) UNIQUE NOT NULL,
  role text NOT NULL REFERENCES roles(value),
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE coordinates (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  entity_id uuid NOT NULL,
  entity_type varchar(20) NOT NULL CHECK (entity_type IN ('driver','passenger')),
  latitude decimal(10,8) NOT NULL,
  longitude decimal(11,8) NOT NULL,
  is_current boolean DEFAULT true
);

CREATE TABLE ride_status (value text PRIMARY KEY);
INSERT INTO ride_status (value) VALUES ('REQUESTED'),('MATCHED'),('IN_PROGRESS'),('COMPLETED'),('CANCELLED');

CREATE TABLE rides (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  ride_number varchar(50) UNIQUE NOT NULL,
  passenger_id uuid NOT NULL REFERENCES users(id),
  driver_id uuid REFERENCES users(id),
  status text REFERENCES ride_status(value) NOT NULL DEFAULT 'REQUESTED',
  requested_at timestamptz DEFAULT now(),
  estimated_fare numeric(10,2)
);

COMMIT;
