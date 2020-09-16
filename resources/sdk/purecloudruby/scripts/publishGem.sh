BUILD_DIR=$1
GEM_NAME=$2
GEM_KEY=$3
INSTALL_RAKE=$4
IS_NEW_RELEASE=$5
VERSION=$6

echo "BUILD_DIR=$BUILD_DIR"
echo "GEM_NAME=$GEM_NAME"
echo "INSTALL_RAKE=$INSTALL_RAKE"
echo "IS_NEW_RELEASE=$IS_NEW_RELEASE"
echo "VERSION=$VERSION"

if [ ! "$IS_NEW_RELEASE" = "true" ]
then
	echo "Skipping gem publish"
	exit 0
fi

GEM_CREDENTIALS_FILE=~/.gem/credentials
GEM_KEY_NAME="developer_evangelists"

cd $BUILD_DIR

# Write files
echo "require 'rubygems'
require 'gems'
task :release do
    Gems.configure do |config|
      config.key = ENV['GEM_KEY']
    end

    puts Gems.push File.new \"$GEM_NAME-$VERSION.gem\"
end" > rakefile

echo "source 'https://rubygems.org'
gem 'gems'" > Gemfile

if [ "$INSTALL_RAKE" == "true" ]
then
	export PATH=$PATH:/home/jenkins/bin

	gem install io-console -v 0.5.4
	gem install rake
	gem env
fi

# Install gems
#gem update --system
echo "Bundle location"
bundle install

# Publish gem
echo "Publishing gems"
gem uninstall psych
rake release --trace
