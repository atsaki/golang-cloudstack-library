package cloudstack

import "fmt"

func (c *Client) GetVirtualMachine(id string) (*VirtualMachine, error) {
	p := NewListVirtualMachinesParameter()
	p.Id.Set(id)
	rs, err := c.ListVirtualMachines(p)
	if err != nil {
		return nil, err
	}
	if len(rs) != 1 {
		return nil, fmt.Errorf("Expected just 1 item, %d items are returned", len(rs))
	}
	return rs[0], nil
}

func (r *VirtualMachine) Refresh() (err error) {
	newr, err := r.client.GetVirtualMachine(r.Id.String())
	if newr != nil {
		*r = *newr
	} else {
		newr = new(VirtualMachine)
		newr.setClient(r.client)
		*r = *newr
	}
	*r = *newr
	return err
}

func (r *VirtualMachine) Update(args map[string]interface{}) (err error) {
	p := NewUpdateVirtualMachineParameter(r.Id.String())
	newr, err := r.client.UpdateVirtualMachine(p)
	if newr != nil {
		*r = *newr
	} else {
		newr = new(VirtualMachine)
		newr.setClient(r.client)
		*r = *newr
	}
	return err
}

func (r *VirtualMachine) Delete() (err error) {
	p := NewDestroyVirtualMachineParameter(r.Id.String())
	p.Expunge.Set(true)
	newr, err := r.client.DestroyVirtualMachine(p)
	if newr != nil {
		*r = *newr
	} else {
		newr = new(VirtualMachine)
		newr.setClient(r.client)
		*r = *newr
	}
	return err
}
