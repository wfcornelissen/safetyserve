# Project Plan: SafeServe

## Project Overview
**Project Name:** SafeServe  
**Purpose:** A cloud-based platform designed to store and manage Health, Safety, and Environment (HSE) documents for companies.

## Use Cases
SafeServe is intended for two types of clients:
1. **Contractors** – Businesses or individuals that need to share HSE compliance documents.
2. **Hiring Companies (Mines)** – Organizations that need to track and manage contractor compliance.

## Core Features
### Contractor Portal
- Dashboard displaying:
  - Upcoming expiry dates of documents.
  - List of mines with access to their documents.
  - Option to share or revoke access to a mine.
- Document Management:
  - Upload new documents.
  - View currently uploaded documents.
  - Remove existing documents.

### Hiring Company (Mine) Portal
- Dashboard displaying:
  - List of contractor documents that have been shared with them.
  - Compliance tracking of contractors.
  - Upcoming document expiry dates.
- Document Management:
  - Request specific document updates from contractors.
  - Set deadlines for updates.
- Compliance Reports:
  - Generate reports on contractor compliance.

## Future Enhancements
- **Bookings System** – Enable contractors to book work schedules.
- **Invoicing** – Allow invoices to be uploaded and managed.
- **Advanced Reporting** – Generate detailed reports on completed work and compliance status.

## Technology Stack
- **Database:** Non-SQL database
- **Backend:** Go (Golang) for middleware server
- **Frontend:** JavaScript (Potential frameworks: React, Vue, or Angular)

## Development Roadmap
### Phase 1: Core Platform
- [ ] Set up non-SQL database.
- [ ] Develop Go-based backend middleware.
- [ ] Implement authentication and user roles.
- [ ] Build contractor and mine dashboards.
- [ ] Develop document upload, viewing, and removal features.
- [ ] Implement access control for document sharing.

### Phase 2: Compliance & Reporting
- [ ] Implement compliance tracking for mines.
- [ ] Add expiry date tracking and notifications.
- [ ] Develop compliance report generation.

### Phase 3: Enhancements
- [ ] Introduce booking system for scheduling work.
- [ ] Implement invoice management features.
- [ ] Expand reporting capabilities.

## Deployment & Hosting
- **Backend Hosting:** Cloud-based Go server.
- **Frontend Hosting:** Web application hosted on cloud services.
- **Database Hosting:** Managed non-SQL database service.

## Conclusion
SafeServe aims to streamline HSE document management for contractors and hiring companies, ensuring compliance and accessibility in a user-friendly manner. With planned enhancements, the platform will provide even greater functionality and value for its users.

## Phases
- [ ] Phase 1: Core Platform
- [ ] Phase 2: Compliance & Reporting
- [ ] Phase 3: Enhancements

## Phase1
- [ ] Develop document upload, viewing, and removal features.
- [ ] Develop Go-based backend middleware.
- [ ] Set up non-SQL database.
- [ ] Implement authentication and user roles.
- [ ] Build contractor and mine dashboards.
- [ ] Implement access control for document sharing.

## Phase2
- [ ] Implement compliance tracking for mines.
- [ ] Add expiry date tracking and notifications.
- [ ] Develop compliance report generation.

## Phase3
- [ ] Introduce booking system for scheduling work.
- [ ] Implement invoice management features.
- [ ] Expand reporting capabilities.
