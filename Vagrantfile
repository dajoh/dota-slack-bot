VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "box-cutter/centos65"
  config.vm.synced_folder ".", "/go/src/github.com/dajoh/dota-slack-bot"
  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "provision/ansible.yml"
    ansible.sudo = true
  end
end
