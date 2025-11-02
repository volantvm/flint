# Flint Development Tracker

**Last Updated**: 2025-11-02
**Version Target**: 1.28.0+

## Overview
This document tracks all GitHub issues and roadmap items for the Flint project. Progress is updated regularly as tasks are completed.

## Priority Matrix

### 🔴 Critical (Blocking Users) - ALL COMPLETED ✅
- ✅ #32, #23 - Cannot create VM from ISO (CloudInit nil error)
- ✅ #30 - Unable to create new storage pools
- ✅ #31 - Installation script broken

### 🟡 High Priority (Usability Issues) - ALL COMPLETED ✅
- ✅ #25 - Storage calculation error (>100% usage)
- ✅ #20 - Missing libvirt-lxc.so.0 dependency
- ✅ #24 - libvirt version compatibility errors

### 🟢 Medium Priority (Platform Support)
- #28 - musl/Alpine Linux support (pending)

### 🔵 Features (Roadmap v1.28.0)
- #19 - noVNC console integration ⭐ ROADMAP
- #17 - PXE installation support ⭐ ROADMAP
- #15 - Firewall feature via libvirt ⭐ ROADMAP

### 🟣 Features (Future)
- #29 - libvirt over SSH support
- #26 - Multi-server management

---

## Bug Fixes

### #32 & #23 - VM Creation from ISO Fails (CloudInit nil)
**Status**: ✅ COMPLETED
**Priority**: CRITICAL
**Affected Users**: Multiple reports
**Root Cause**: Frontend sending wrong field name for ISO installations

**Issue Details**:
- Creating VMs from ISO images fails silently
- Console shows: "DEBUG: CloudInit is nil"
- "Create & Start VM" button does nothing
- Affects both Oracle 9.6 and other distros

**Root Cause Identified**:
1. Frontend sent `ISOPath` field for ISO installations
2. Backend VMCreationConfig struct only has `ImageName` field (no `ISOPath`)
3. Validation requires `ImageName` to be non-empty
4. For ISO installs, `ImageName` was empty → validation failed silently

**Solution Implemented**:
1. [x] Updated `simple-vm-wizard.tsx` to send `imageName` for both ISO and template types
2. [x] Removed unused `ISOPath` field from `create-vm-wizard.tsx`
3. [x] Fixed JSON field name casing (`cloudInit`, `commonFields`, `rawUserData`)
4. [x] Added `enableCloudInit` field to form submission

**Files Modified**:
- `/web/components/simple-vm-wizard.tsx`
- `/web/components/create-vm-wizard.tsx`

**Commit**: 06c1294

---

### #31 - Installation Script Broken (Repository Migration)
**Status**: ✅ COMPLETED
**Priority**: CRITICAL
**Affected Users**: New installations

**Issue Details**:
- Install script uses old repository: `ccheshirecat/flint`
- Should use: `volantvm/flint`
- Script fails with "Moved Permanently" API redirect
- GitHub API returns 301 redirect instead of following it

**Solution Implemented**:
1. [x] Updated `/install.sh` REPO variable to `volantvm/flint`
2. [x] Updated `/install-systemd.sh` documentation references
3. [x] Updated README.md badges and install instructions
4. [x] Updated docs.md install command
5. [x] Migrated Go module path from `github.com/ccheshirecat/flint` to `github.com/volantvm/flint`
6. [x] Updated all Go import statements across 20+ files

**Files Modified**:
- `/install.sh`
- `/install-systemd.sh`
- `/README.md`
- `/docs.md`
- `/go.mod`
- All Go source files with imports

**Commit**: 64454d9

---

### #30 - Storage Pool Creation Not Working
**Status**: ✅ COMPLETED
**Priority**: CRITICAL
**Affected Users**: Debian 13, Devuan 6

**Issue Details**:
- "Add Pool" button click does nothing
- Tested with root and libvirt-qemu group users
- Tried various storage locations and permissions
- Affects both directory and filesystem pool types

**Root Cause Identified**:
Storage pool creation was completely unimplemented across all layers:
1. No backend libvirtclient function to create pools
2. No HTTP handler for pool creation
3. No HTTP route (POST /storage-pools)
4. No frontend API method
5. Frontend form had no state management or submit handler

**Solution Implemented**:
1. [x] Added `PoolConfig` type to `/pkg/core/types.go`
2. [x] Implemented `CreateStoragePool()` in `/pkg/libvirtclient/storage.go`
3. [x] Added `handleCreateStoragePool()` in `/server/handlers.go`
4. [x] Added `POST /storage-pools` route in `/server/server.go`
5. [x] Added `createPool()` API method in `/web/lib/api.ts`
6. [x] Implemented complete form with state management in `/web/components/storage-view.tsx`
7. [x] Added validation, error handling, and toast notifications

**Files Modified**:
- `/pkg/core/types.go`
- `/pkg/libvirtclient/storage.go`
- `/server/handlers.go`
- `/server/server.go`
- `/web/lib/api.ts`
- `/web/components/storage-view.tsx`

**Commit**: 33355d5

---

### #25 - Storage Calculation Error (>100% Usage)
**Status**: 🟡 NOT STARTED
**Priority**: HIGH

**Issue Details**:
- Dashboard shows "109% storage usage"
- Reports 2042 GB used / 187 GB total
- Storage page shows 1.8 TB total (conflicting)
- Multiple pools on same partition cause double-counting

**Investigation Plan**:
1. [ ] Find storage calculation logic
2. [ ] Review how pool capacities are aggregated
3. [ ] Fix double-counting when pools share underlying filesystem
4. [ ] Ensure consistent reporting across dashboard and storage pages
5. [ ] Add tests for overlapping pool scenarios

**Files to Examine**:
- `/pkg/libvirtclient/storage.go`
- Dashboard handlers in `/server/handlers/`
- Frontend dashboard components

**Solution Approach**: Deduplicate storage calculations, track underlying filesystems

---

### #24 - libvirt Version Compatibility
**Status**: ✅ COMPLETED
**Priority**: HIGH

**Issue Details**:
- Flint requires libvirt >= 6.10.0
- Many systems (OpenEuler 22.03 SP4) ship with 6.2
- Error: "version LIBVIRT_6.10.0 not found"
- No documentation of minimum requirements

**Solution Implemented**:
1. [x] Added comprehensive Prerequisites section to README.md with:
   - System requirements (libvirt >= 6.10.0, QEMU/KVM)
   - Distro-specific installation instructions (Debian/Ubuntu, RHEL/Fedora, Arch)
2. [x] Added runtime version check in pkg/libvirtclient/client.go:
   - Checks version immediately after libvirt connection
   - Displays helpful error with current version and minimum required
   - Directs users to README for upgrade instructions
3. [x] Version check prevents confusing errors from unsupported features

**Files Modified**:
- `/README.md` - Added prerequisites section
- `/pkg/libvirtclient/client.go` - Added version check in NewClient()

**Commits**: 01f27a1, 7ee28a7

---

### #20 - Missing libvirt-lxc.so.0 Dependency
**Status**: ✅ COMPLETED
**Priority**: HIGH

**Issue Details**:
- Binary requires libvirt-lxc.so.0 shared library
- Not always installed by default
- Prevents flint from starting

**Solution Implemented**:
1. [x] Added note to README.md prerequisites section with solution:
   - Documented the libvirt-lxc.so.0 error users may encounter
   - Provided installation commands for all major distros:
     - Debian/Ubuntu: `sudo apt install -y libvirt-daemon-driver-lxc`
     - RHEL/Fedora: `sudo dnf install -y libvirt-daemon-lxc`
     - Arch: `sudo pacman -S libvirt-lxc`
2. [x] Users can now self-resolve this issue using README

**Files Modified**:
- `/README.md` - Added troubleshooting note in prerequisites

**Commit**: 01f27a1

**Note**: Static linking not pursued as it would significantly increase binary size and complexity

---

### #28 - musl/Alpine Linux Support
**Status**: 🟢 NOT STARTED
**Priority**: MEDIUM

**Issue Details**:
- Binary requires glibc
- Doesn't work on musl-based systems (Alpine Linux)
- Request for musl variant or static binary

**Investigation Plan**:
1. [ ] Research libvirt-go compatibility with musl
2. [ ] Add musl build target to Makefile
3. [ ] Test musl binary on Alpine
4. [ ] Update GitHub Actions for musl releases
5. [ ] Consider fully static binary (may be complex with CGO)

**Files to Update**:
- `/Makefile`
- `/.github/workflows/release.yml`

**Solution Approach**: Add musl build target if libvirt-go supports it

---

## Feature Implementations (ROADMAP v1.28.0)

### #19 - noVNC Console Integration ⭐
**Status**: 🔵 NOT STARTED
**Priority**: ROADMAP FEATURE
**Roadmap**: v1.28.0

**Feature Description**:
- Web-based VNC console access to VMs
- Use noVNC library for browser-based access
- No client software required

**Implementation Plan**:
1. [ ] Research noVNC integration patterns
2. [ ] Add noVNC library to frontend
3. [ ] Create WebSocket proxy for VNC traffic
4. [ ] Add VNC console UI component
5. [ ] Integrate with VM detail page
6. [ ] Handle VNC authentication and security
7. [ ] Test with various VM configurations

**Technical Requirements**:
- noVNC JavaScript library
- WebSocket server for VNC proxy
- VNC configuration in libvirt domains
- Security considerations (authentication, encryption)

**Files to Create/Modify**:
- Frontend: `/web/src/components/VNCConsole.tsx`
- Backend: `/server/handlers/vnc.go`
- WebSocket: `/pkg/vnc/proxy.go`

**Dependencies**: None (standalone feature)

---

### #17 - PXE Installation Support ⭐
**Status**: 🔵 NOT STARTED
**Priority**: ROADMAP FEATURE
**Roadmap**: v1.28.0

**Feature Description**:
- Allow VM installation via PXE boot
- Support network-based OS installation
- Configure PXE boot parameters

**Implementation Plan**:
1. [ ] Research libvirt PXE boot configuration
2. [ ] Design UI for PXE boot options
3. [ ] Add PXE network configuration to VM creation
4. [ ] Support kernel/initrd URLs
5. [ ] Add boot parameters input
6. [ ] Implement virt-install style PXE flow
7. [ ] Test with common PXE servers

**Technical Requirements**:
- Libvirt network boot XML configuration
- PXE server URL inputs
- Kernel parameters configuration
- Network bridge/NAT setup

**Example Command** (from issue):
```bash
virt-install --name=vm1.name --memory=4096 --vcpus=4 \
--disk /dev/vg0/vm1-name-root,bus=virtio,cache=none \
--os-variant ol8.10 -x 'inst.ks=http://url_path/ks/ks-vm1-name.cfg' \
--network bridge=br10 --graphics none --virt-type kvm \
-l http://url_path/OL8_aarch64/ --autostart
```

**Files to Create/Modify**:
- Frontend: `/web/src/components/VMCreation/*`
- Backend: `/pkg/libvirtclient/vm_create.go`
- Models: Add PXE configuration structs

**Dependencies**: None

---

### #15 - Firewall Feature via libvirt ⭐
**Status**: 🔵 NOT STARTED
**Priority**: ROADMAP FEATURE
**Roadmap**: v1.28.0

**Feature Description**:
- Manage VM firewall rules via libvirt
- Configure nwfilter rules
- Network traffic filtering per VM

**Implementation Plan**:
1. [ ] Research libvirt nwfilter capabilities
2. [ ] Design firewall rule UI/UX
3. [ ] Implement nwfilter management in libvirtclient
4. [ ] Create firewall rule CRUD operations
5. [ ] Add firewall configuration to VM settings
6. [ ] Support common rule types (IP, port, protocol)
7. [ ] Test firewall rule enforcement

**Technical Requirements**:
- libvirt nwfilter XML manipulation
- Firewall rule editor UI
- Rule validation and testing
- Integration with VM network interfaces

**Files to Create/Modify**:
- Frontend: `/web/src/components/Firewall/*`
- Backend: `/pkg/libvirtclient/firewall.go`
- Backend: `/server/handlers/firewall.go`

**Dependencies**: None

---

## Future Features

### #29 - libvirt over SSH Support
**Status**: 🟣 NOT STARTED
**Priority**: FUTURE FEATURE

**Feature Description**:
- Connect to remote libvirt servers via SSH
- Support URI format: `qemu+ssh://user@host/system`
- Manage distributed KVM/QEMU setups

**Implementation Plan**:
1. [ ] Add SSH connection configuration UI
2. [ ] Update libvirt connection initialization
3. [ ] Support connection URI input
4. [ ] Handle SSH authentication (keys, passwords)
5. [ ] Test remote server management
6. [ ] Add connection health monitoring

**Files to Create/Modify**:
- `/pkg/libvirtclient/connection.go`
- Frontend: Connection settings component
- Backend: SSH credential management

**Dependencies**: Related to #26 (multi-server management)

---

### #26 - Multi-Server Management
**Status**: 🟣 NOT STARTED
**Priority**: FUTURE FEATURE

**Feature Description**:
- Manage multiple VM servers from single Flint instance
- Switch between different hosts
- Aggregate resource monitoring

**Implementation Plan**:
1. [ ] Design multi-server architecture
2. [ ] Create server registry/configuration
3. [ ] Add server switching UI
4. [ ] Update all handlers to support server context
5. [ ] Implement per-server resource monitoring
6. [ ] Test with multiple libvirt hosts

**Files to Create/Modify**:
- Database: Server configuration storage
- Frontend: Server selector component
- Backend: Connection pool management

**Dependencies**: May benefit from #29 (SSH support)

---

## Progress Summary

### Overall Statistics
- **Total Issues**: 13
- **Bugs**: 8
- **Features**: 5
- **Completed**: 5 ✅
- **In Progress**: 0
- **Not Started**: 8

### By Priority
- 🔴 Critical: 0 remaining (3/3 completed) ✅
- 🟡 High: 1 remaining (2/3 completed)
- 🟢 Medium: 1 issue
- 🔵 Roadmap Features: 3 features
- 🟣 Future Features: 2 features

### Roadmap v1.28.0 Progress
- [ ] PXE install flow (#17)
- [ ] noVNC integration (#19)
- [ ] Firewall feature (#15)

**Roadmap Completion**: 0/3 (0%)

---

## Development Workflow

### Bug Fix Workflow
1. Investigate issue and identify root cause
2. Create feature branch: `fix/issue-NUMBER-description`
3. Implement fix with tests
4. Update this tracker
5. Create PR with issue reference
6. Close issue after merge

### Feature Development Workflow
1. Design feature architecture
2. Create feature branch: `feature/issue-NUMBER-description`
3. Implement feature incrementally
4. Add tests and documentation
5. Update this tracker
6. Create PR with issue reference
7. Close issue after merge

### Testing Requirements
- Unit tests for business logic
- Integration tests for API endpoints
- Manual testing on target platforms
- Update test documentation

---

## Notes & Blockers

### Current Blockers
None identified yet.

### Technical Debt
- Need comprehensive error handling across codebase
- API endpoints need consistent error responses
- Frontend needs better error user feedback
- Missing integration tests for VM lifecycle

### Future Considerations
- Multi-tenancy support
- RBAC (Role-Based Access Control)
- Audit logging
- Backup/restore functionality
- VM migration tools
- Cluster support

---

## Update Log

### 2025-11-02
- Created initial development tracker
- Catalogued all 13 GitHub issues
- Analyzed roadmap v1.28.0 requirements
- Explored codebase structure
- Identified priority matrix
- ✅ Fixed #31: Repository migration (install scripts + Go module)
- ✅ Fixed #32/#23: VM creation from ISO (frontend field mapping)
- ✅ Fixed #30: Storage pool creation (full stack implementation)
- ✅ Fixed #25: Storage calculation errors (deduplicated filesystem counting)
- ✅ Fixed #24: libvirt version requirements (README + runtime check)
- ✅ Fixed #20: libvirt-lxc dependency (documented in prerequisites)
- Fixed docs.md link in README (lowercase)
- **All critical and high-priority bugs resolved** 🎉
- Next: Roadmap v1.28.0 features (#19 noVNC, #17 PXE, #15 Firewall)
