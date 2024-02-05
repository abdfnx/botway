# frozen_string_literal: true

require_relative "lib/bwrb/version"

Gem::Specification.new do |spec|
  spec.name          = "bwrb"
  spec.version       = Botwayrb::VERSION
  spec.authors       = ["abdfnx"]

  spec.summary       = "Ruby client package for Botway."
  spec.description   = spec.summary
  spec.homepage      = "https://github.com/abdfnx/botway"
  spec.license       = "MIT"
  spec.required_ruby_version = Gem::Requirement.new(">= 2.4.0")

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = "https://rubygems.org/gems/bwrb"
  spec.metadata["changelog_uri"] = "https://github.com/abdfnx/botway/main/CHANGELOG.md"

  spec.files = Dir.chdir(File.expand_path(__dir__)) do
    `git ls-files -z`.split("\x0").reject { |f| f.match(%r{\A(?:test|spec|features)/}) }
  end

  spec.bindir        = "exe"
  spec.executables   = spec.files.grep(%r{\Aexe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]

  spec.add_dependency "yaml", ">= 0.2", "< 0.4"
  spec.add_dependency "json", "~> 2.6"

  spec.add_development_dependency "bundler", ">= 1.10", "< 3"
  spec.add_development_dependency "rake", "~> 13.0"
  spec.add_development_dependency "rspec", "~> 3.13.0"
  spec.add_development_dependency "rspec-prof", "~> 0.0.7"
  spec.add_development_dependency "rubocop", "~> 1.60.2"
  spec.add_development_dependency "rubocop-performance", "~> 1.0"
  spec.add_development_dependency "rubocop-rake", "~> 0.6.0"
end
