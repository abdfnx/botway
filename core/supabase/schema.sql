CREATE TABLE projects (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES auth.users ON DELETE CASCADE,
  name TEXT NOT NULL,
  repo TEXT NOT NULL,
  platform TEXT NOT NULL,
  lang TEXT NOT NULL,
  package_manager TEXT NOT NULL,
  icon TEXT NULL,
  root_directory TEXT NULL,
  bot_token CHARACTER varying NOT NULL,
  bot_app_token CHARACTER varying NULL,
  bot_secret_token CHARACTER varying NULL,
  zeabur_project_id CHARACTER varying NOT NULL,
  zeabur_service_id CHARACTER varying NOT NULL,
  zeabur_env_id CHARACTER varying NOT NULL,
  build_command CHARACTER varying NULL,
  start_command CHARACTER varying NULL,
  enable_ce BOOLEAN NOT NULL DEFAULT false,
  ce_service_id CHARACTER varying NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, NOW()) NOT NULL,
  CONSTRAINT projects_zeabur_project_id_key UNIQUE (zeabur_project_id)
);

-- Enable row level security
ALTER TABLE projects ENABLE ROW LEVEL SECURITY;

-- Create policies to restrict access to the projects table
CREATE POLICY "Users can insert their own projects."
  ON projects FOR INSERT
  WITH CHECK ( auth.uid() = user_id );

CREATE POLICY "Users can select their own projects."
  ON projects FOR SELECT
  USING ( auth.uid() = user_id );

CREATE POLICY "Users can update their own projects."
  ON projects FOR UPDATE
  USING ( auth.uid() = user_id );

CREATE POLICY "Users can delete their own projects."
  ON projects FOR DELETE
  USING ( auth.uid() = user_id );
