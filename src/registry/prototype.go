package registry

import (
	adminController "github.com/pudthaiiii/go-ibooking/src/app/http/admin/controllers/prototype"
	adminService "github.com/pudthaiiii/go-ibooking/src/app/services/admin"
)

// apply service to controller
func (r *registry) NewPrototypeController() adminController.PrototypeController {
	return adminController.NewPrototypeController(r.AdminService())
}

// apply repository to service
func (r *registry) AdminService() adminService.PrototypeService {
	return adminService.NewPrototypeService(r.NewUsageItemRepository(), r.NewProductRepository())
}