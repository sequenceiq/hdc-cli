package cli

import (
	"errors"
	swagerrors "github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

func (s *ClusterSkeleton) Validate() error {
	var res []error
	if err := validate.RequiredString("ClusterName", "body", string(s.ClusterName)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("HDPVersion", "body", string(s.HDPVersion)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("ClusterType", "body", string(s.ClusterType)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredNumber("InstanceCount", "worker", float64(s.Worker.InstanceCount)); err != nil {
		res = append(res, err)
	} else if s.Worker.InstanceCount < 1 {
		res = append(res, swagerrors.New(1, "The instance count has to be greater than 0"))
	}
	if err := validate.RequiredString("SSHKeyName", "body", string(s.SSHKeyName)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("RemoteAccess", "body", string(s.RemoteAccess)); err != nil {
		res = append(res, err)
	}
	if err := validate.Required("WebAccess", "body", s.WebAccess); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("ClusterAndAmbariUser", "body", string(s.ClusterAndAmbariUser)); err != nil {
		res = append(res, err)
	}
	if err := validate.RequiredString("ClusterAndAmbariPassword", "body", string(s.ClusterAndAmbariPassword)); err != nil {
		res = append(res, err)
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			for _, e := range err {
				res = append(res, e)
			}
		}
	}
	if s.HiveMetastore != nil {
		if err := s.HiveMetastore.Validate(); err != nil {
			for _, e := range err {
				res = append(res, e)
			}
		}
	}

	if len(res) > 0 {
		return swagerrors.CompositeValidationError(res...)
	}
	return nil
}

func (n *Network) Validate() []error {
	var res []error = nil

	if len(n.VpcId) > 0 || len(n.SubnetId) > 0 {
		if err := validate.RequiredString("VpcId", "network", n.VpcId); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("SubnetId", "network", n.SubnetId); err != nil {
			res = append(res, err)
		}
	}

	return res
}

func (h *HiveMetastore) Validate() []error {
	var res []error = nil

	if len(h.DatabaseType) > 0 || len(h.Username) > 0 || len(h.Password) > 0 || len(h.URL) > 0 {
		if err := validate.RequiredString("Name", "hivemetastore", h.Name); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("DatabaseType", "hivemetastore", h.DatabaseType); err != nil {
			res = append(res, err)
		} else if h.DatabaseType != POSTGRES || h.DatabaseType != MYSQL {
			res = append(res, errors.New("Invalid database type. Accepted values are: POSTGRES and MYSQL."))
		}
		if err := validate.RequiredString("Password", "hivemetastore", h.Password); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("Username", "hivemetastore", h.Username); err != nil {
			res = append(res, err)
		}
		if err := validate.RequiredString("URL", "hivemetastore", h.URL); err != nil {
			res = append(res, err)
		}
	}

	return res
}
