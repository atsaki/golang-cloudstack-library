package cloudstack

import "fmt"

func (vm *VirtualMachine) Refresh() (Resource, error) {
	p := NewListVirtualMachinesParameter()
	p.Id.Set(vm.Id)
	vms, err := vm.Client.ListVirtualMachines(p)
	if err != nil {
		return nil, err
	}
	if len(vms) != 1 {
		return nil, fmt.Errorf("Expected just 1 item, %d items are returned", len(vms))
	}
	return vms[0], nil
}

func (vm *VirtualMachine) Update(args map[string]interface{}) (Resource, error) {
	return nil, nil
}

func (vm *VirtualMachine) Delete() error {
	p := NewDestroyVirtualMachineParameter(vm.Id.String())
	p.Expunge.Set(true)
	_, err := vm.Client.DestroyVirtualMachine(p)
	return err
}
