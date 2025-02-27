package server

import (
	"net/http"

	"github.com/containers/podman/v4/pkg/api/handlers/libpod"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerManifestHandlers(r *mux.Router) error {
	v3 := r.PathPrefix("/v{version:[0-3][0-9A-Za-z.-]*}/libpod/manifests").Subrouter()
	v4 := r.PathPrefix("/v{version:[4-9][0-9A-Za-z.-]*}/libpod/manifests").Subrouter()
	// swagger:operation POST /libpod/manifests/{name}/push manifests ManifestPushV3Libpod
	// ---
	// summary: Push manifest to registry
	// description: |
	//   Push a manifest list or image index to a registry
	//
	//   Deprecated: As of 4.0.0 use ManifestPushLibpod instead
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the manifest
	//  - in: query
	//    name: destination
	//    type: string
	//    required: true
	//    description: the destination for the manifest
	//  - in: query
	//    name: all
	//    description: push all images
	//    type: boolean
	// responses:
	//   200:
	//     schema:
	//       $ref: "#/definitions/IDResponse"
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   500:
	//     $ref: "#/responses/InternalError"
	v3.Handle("/{name}/push", s.APIHandler(libpod.ManifestPushV3)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/manifests/{name}/registry/{destination} manifests ManifestPushLibpod
	// ---
	// summary: Push manifest list to registry
	// description: |
	//   Push a manifest list or image index to the named registry
	//
	//   As of v4.0.0
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the manifest list
	//  - in: path
	//    name: destination
	//    type: string
	//    required: true
	//    description: the registry for the manifest list
	//  - in: query
	//    name: all
	//    description: push all images
	//    type: boolean
	//    default: false
	//  - in: query
	//    name: tlsVerify
	//    type: boolean
	//    default: false
	//    description: skip TLS verification for registries
	// responses:
	//   200:
	//     schema:
	//       $ref: "#/definitions/IDResponse"
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   500:
	//     $ref: "#/responses/InternalError"
	v4.Handle("/{name:.*}/registry/{destination:.*}", s.APIHandler(libpod.ManifestPush)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/manifests manifests ManifestCreateLibpod
	// ---
	// summary: Create
	// description: Create a manifest list
	// produces:
	// - application/json
	// parameters:
	// - in: query
	//   name: name
	//   type: string
	//   description: manifest list or index name to create
	//   required: true
	// - in: query
	//   name: images
	//   type: string
	//   required: true
	//   description: |
	//     One or more names of an image or a manifest list. Repeat parameter as needed.
	//
	//     Support for multiple images, as of version 4.0.0
	//     Alias of `image` is support for compatibility with < 4.0.0
	//     Response status code is 200 with < 4.0.0 for compatibility
	// - in: query
	//   name: all
	//   type: boolean
	//   description: add all contents if given list
	// - in: body
	//   name: options
	//   description: options for new manifest
	//   required: false
	//   schema:
	//     $ref: "#/definitions/ManifestModifyOptions"
	// responses:
	//   201:
	//     schema:
	//       $ref: "#/definitions/IDResponse"
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   404:
	//     $ref: "#/responses/NoSuchImage"
	//   500:
	//     $ref: "#/responses/InternalError"
	v3.Handle("/create", s.APIHandler(libpod.ManifestCreate)).Methods(http.MethodPost)
	v4.Handle("/{name:.*}", s.APIHandler(libpod.ManifestCreate)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/manifests/{name}/exists manifests ManifestExistsLibpod
	// ---
	// summary: Exists
	// description: |
	//   Check if manifest list exists
	//
	//   Note: There is no contract that the manifest list will exist for a follow-on operation
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the manifest list
	// produces:
	// - application/json
	// responses:
	//   204:
	//     description: manifest list exists
	//   404:
	//     $ref: '#/responses/NoSuchManifest'
	//   500:
	//     $ref: '#/responses/InternalError'
	v3.Handle("/{name:.*}/exists", s.APIHandler(libpod.ManifestExists)).Methods(http.MethodGet)
	v4.Handle("/{name:.*}/exists", s.APIHandler(libpod.ManifestExists)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/manifests/{name}/json manifests ManifestInspectLibpod
	// ---
	// summary: Inspect
	// description: Display attributes of given manifest list
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the manifest list
	// responses:
	//   200:
	//     $ref: "#/responses/InspectManifest"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   500:
	//     $ref: "#/responses/InternalError"
	v3.Handle("/{name:.*}/json", s.APIHandler(libpod.ManifestInspect)).Methods(http.MethodGet)
	v4.Handle("/{name:.*}/json", s.APIHandler(libpod.ManifestInspect)).Methods(http.MethodGet)
	// swagger:operation PUT /libpod/manifests/{name} manifests ManifestModifyLibpod
	// ---
	// summary: Modify manifest list
	// description: |
	//   Add/Remove an image(s) to a manifest list
	//
	//   Note: operations are not atomic when multiple Images are provided.
	//
	//   As of v4.0.0
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the manifest
	//  - in: query
	//    name: tlsVerify
	//    type: boolean
	//    default: false
	//    description: skip TLS verification for registries
	//  - in: body
	//    name: options
	//    description: options for mutating a manifest
	//    required: true
	//    schema:
	//        $ref: "#/definitions/ManifestModifyOptions"
	// responses:
	//   200:
	//     schema:
	//       $ref: "#/definitions/ManifestModifyReport"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   409:
	//     description: Operation had partial success, both Images and Errors may have members
	//     schema:
	//       $ref: "#/definitions/ManifestModifyReport"
	//   500:
	//     $ref: "#/responses/InternalError"
	v4.Handle("/{name:.*}", s.APIHandler(libpod.ManifestModify)).Methods(http.MethodPut)
	// swagger:operation POST /libpod/manifests/{name}/add manifests ManifestAddLibpod
	// ---
	// summary: Add image
	// description: |
	//   Add an image to a manifest list
	//
	//   Deprecated: As of 4.0.0 use ManifestModifyLibpod instead
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the manifest
	//  - in: body
	//    name: options
	//    description: options for creating a manifest
	//    schema:
	//      $ref: "#/definitions/ManifestAddOptions"
	// responses:
	//   200:
	//     schema:
	//       $ref: "#/definitions/IDResponse"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   409:
	//     $ref: "#/responses/BadParamError"
	//   500:
	//     $ref: "#/responses/InternalError"
	v3.Handle("/{name:.*}/add", s.APIHandler(libpod.ManifestAddV3)).Methods(http.MethodPost)
	// swagger:operation DELETE /libpod/manifests/{name} manifests ManifestDeleteV3Libpod
	// ---
	// summary: Remove image from a manifest list
	// description: |
	//   Remove an image from a manifest list
	//
	//   Deprecated: As of 4.0.0 use ManifestModifyLibpod instead
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the image associated with the manifest
	//  - in: query
	//    name: digest
	//    type: string
	//    description: image digest to be removed
	// responses:
	//   200:
	//     schema:
	//       $ref: "#/definitions/IDResponse"
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   500:
	//     $ref: "#/responses/InternalError"
	v3.Handle("/{name:.*}", s.APIHandler(libpod.ManifestRemoveDigestV3)).Methods(http.MethodDelete)
	// swagger:operation DELETE /libpod/manifests/{name} manifests ManifestDeleteLibpod
	// ---
	// summary: Delete manifest list
	// description: |
	//   Delete named manifest list
	//
	//   As of v4.0.0
	// produces:
	// - application/json
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: The name or ID of the  list to be deleted
	// responses:
	//   200:
	//     $ref: "#/responses/DocsLibpodImagesRemoveResponse"
	//   404:
	//     $ref: "#/responses/NoSuchManifest"
	//   500:
	//     $ref: "#/responses/InternalError"
	v4.Handle("/{name:.*}", s.APIHandler(libpod.ManifestDelete)).Methods(http.MethodDelete)
	return nil
}
