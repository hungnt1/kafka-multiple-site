# Apache Zookeeper Ansible

It is group of playbooks to manage apache zookeeper.

## There are 2 ssl in cluster.
01: ssl for kafka to zk and zk to zk
02: ssl for client to kafka and kafka and kafka

## **Requirements**
* netaddr ( Mandatory )
* Download Apache Zookeeper Tar Manually ( Mandatory )
* vagrant ( Optional )
* Any OS with SystemD ( Mandatory )
* Ansible ( Mandatory )

## **Notes***
```
1. Installation for NetAddr Module.
   * pip install netaddr
   * https://docs.ansible.com/ansible/latest/user_guide/playbooks_filters_ipaddr.html

2. All tasks like jvm/logging/downgrade/removeOldVersion will be done in serial order.
```

## **Development Environment Setup**
* **STEP-1**
```
vagrant up
```
* **STEP-2**

add below lines to `/etc/hosts` file in all nodes.

- File host worker
```
192.168.50.190 zookeeper01
192.168.50.191 zookeeper02
192.168.50.192 zookeeper03

192.168.30.190 kafka01
192.168.30.191 kafka02
192.168.30.192 kafka03

```

- File host ansible play
```
192.168.50.190 zookeeper01
192.168.50.191 zookeeper02
192.168.50.192 zookeeper03

192.168.50.190 kafka01
192.168.50.191 kafka02
192.168.50.192 kafka03

```

copy ssh key from manage node to worker node
```
ssh-copy-id hungnt1@zookeeper1
ssh-copy-id hungnt1@zookeeper2
ssh-copy-id hungnt1@zookeeper3

```


* **STEP-3**
```
ansible-playbook -i inventory/development/cluster.ini clusterSetup.yml
```

## **Apache Zookeeper Playbooks**

## **Cloud Infra Using Terraform**

* `terraform/aws`
* `terraform/oci`

## **Cloud Infra Using Pulumi**

* `pulumi/aws`

### **AWS Cloud PreSetup for cluster**

It will enable following things on all nodes.

1. `/zookeeper` mount point from ebs created by terraform.
2. Install and configure `awslogs` agent for kafka-logs.
3. Install python3 packages

* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .
* Update Required vars in ```inventory/<environment>/cluster.ini``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterAwsPreSetup.yml```



### **To start new cluster***
* download zookeeper with version zookeeperVersion: wget https://mirror.downloadvn.com/apache/zookeeper/zookeeper-3.6.3/apache-zookeeper-3.6.3-bin.tar.gz 
* Install addition repo
```
sudo add-apt-repository universe
sudo apt update
```
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .
* Update Required vars in ```inventory/<environment>/cluster.ini``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterSetup.yml```

### **Monitoring Setup**
* **To add custom metric exporter to cluster**

```ansible-playbook -i inventory/<environment>/cluster.ini clusterCustomMetricExporter.yml```

* **To add newrelic monitoring to cluster**

```ansible-playbook -i inventory/<environment>/cluster.ini clusterNewRelicSetup.yml```

### **Rolling restart cluster**

```ansible-playbook -i inventory/<environment>/cluster.ini clusterRollingRestart.yml```

### **To update jvm settings of cluster**
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterJvmConfigs.yml```

### **To update logging settings of cluster**
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterLogging.yml```

### **To upgrade java version of cluster**
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterJava.yml```

### **To upgrade OS version of cluster**
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterSystemUpgrade.yml```

### **To remove old version files of zookeeper from cluster**
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterRemoveOldVersion.yml```

### **To remove zookeeper from cluster**
* Update Required vars in ```inventory/<environment>/group_vars/all.yml``` .

```ansible-playbook -i inventory/<environment>/cluster.ini clusterRemoveNodes.yml```

### **Tested OS**
* CentOS 7
* RedHat 7

### **Tested Ansible Version**
```
ansible: 3.2.0
ansible-base: 2.10.7
```
