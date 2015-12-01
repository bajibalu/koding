package softlayer

import (
	"koding/db/mongodb/modelhelper"
	"koding/kites/kloud/machinestate"

	"golang.org/x/net/context"
)

// Stop stops the given machine
func (m *Machine) Stop(ctx context.Context) error {
	if err := modelhelper.ChangeMachineState(m.Id, "Machine is stopping", machinestate.Stopping); err != nil {
		return err
	}

	//Get the SoftLayer virtual guest service
	svc, err := m.Session.SLClient.GetSoftLayer_Virtual_Guest_Service()
	if err != nil {
		return err
	}

	_, err = svc.PowerOff(m.Meta.Id)
	if err != nil {
		return err
	}

	if err := waitState(svc, m.Meta.Id, "HALTED"); err != nil {
		return err
	}

	return m.MarkAsStoppedWithReason("Machine is stopped")
}
