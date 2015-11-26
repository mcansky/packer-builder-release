# Packer AWS Release builder

This is a builder plugin for Packer.

----

Packer is a tool for building identical machine images for multiple platforms
from a single source configuration.

----

The aim of this builder is to avoid creating an AMI and instead use packer ability to create, use and then destroy an AWS EC2 instance to build releases of an application.

It's based on Packer own EBS builder with a very small twist : the AMI creation has been removed. So the use of this builder is similar to the EBS one but allows to proceed with several commands on the temporary instance without the creation of an AMI.

### Tests

Tests have been adapted and will run with ```make``` or ```make test```.

### Build

```make bin``` will create the binaries needed. Recommend 1.5.x.

