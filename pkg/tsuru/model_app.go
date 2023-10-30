/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Tsuru app.
type App struct {
	// App name.
	Name string `json:"name"`
	// Cluster name
	Cluster string `json:"cluster,omitempty"`
	// CNames of App
	Cname []string `json:"cname,omitempty"`
	// Number of Deploys
	Deploys int64 `json:"deploys,omitempty"`
	// Unit metrics.
	UnitsMetrics []UnitMetrics `json:"unitsMetrics,omitempty"`
	// Autoscale Recommendations
	AutoscaleRecommendation []RecommendedResources `json:"autoscaleRecommendation,omitempty"`
	// Errors during AppGet
	Error string `json:"error,omitempty"`
	Quota Quota  `json:"quota,omitempty"`
	// Service instance binds on the app
	ServiceInstanceBinds []AppServiceInstanceBinds `json:"serviceInstanceBinds,omitempty"`
	Routers              []AppRouters              `json:"routers,omitempty"`
	InternalAddresses    []AppInternalAddresses    `json:"internalAddresses,omitempty"`
	VolumeBinds          []AppVolumeBinds          `json:"volumeBinds,omitempty"`
	// App tags.
	Tags           []string            `json:"tags,omitempty"`
	Metadata       Metadata            `json:"metadata,omitempty"`
	ProcessesTweak []AppProcessesTweak `json:"processesTweak,omitempty"`
	// App router.
	Router string `json:"router,omitempty"`
	// Custom router options.
	Routeropts map[string]string `json:"routeropts,omitempty"`
	Plan       Plan              `json:"plan,omitempty"`
	Lock       AppLock           `json:"lock,omitempty"`
	// App pool.
	Pool string `json:"pool,omitempty"`
	// App provisioner.
	Provisioner string `json:"provisioner,omitempty"`
	// App platform.
	Platform string `json:"platform,omitempty"`
	// App description.
	Description string `json:"description,omitempty"`
	// Team that owns the app.
	TeamOwner string          `json:"teamOwner,omitempty"`
	Teams     []string        `json:"teams,omitempty"`
	Units     []Unit          `json:"units,omitempty"`
	Ip        string          `json:"ip,omitempty"`
	Owner     string          `json:"owner,omitempty"`
	Autoscale []AutoScaleSpec `json:"autoscale,omitempty"`
}
