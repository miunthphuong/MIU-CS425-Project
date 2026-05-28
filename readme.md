# CarShareX

CarShareX is a car rental management system for connecting customers with
vehicle suppliers, including individual car owners and rental companies. The
platform is designed to make car discovery, booking, identity verification,
payment, and supplier management safer and easier for both sides of the rental
marketplace.

The project was created for the MIU CS425 Software Engineering course.

## Problem Statement

Traditional car rental platforms are often built around large rental companies.
That leaves individual car owners with limited ways to rent out their vehicles
and gives customers fewer choices, higher prices, and less flexible search
options.

CarShareX addresses this by providing a centralized marketplace where customers
can search for cars by location, rental dates, price, car type, and supplier
rating. The system also focuses on trust and safety through supplier
verification, customer identity checks, driving license validation, payment
authentication, and face recognition.

## Main Users

- Customers search for cars, verify their identity, book vehicles, make
  payments, manage bookings, and leave reviews.
- Suppliers add and manage cars, respond to booking requests, track rentals,
  and view reports.
- Admins verify suppliers, manage users, monitor payments, generate reports,
  and resolve disputes.

## Core Features

- Search available cars by pickup location, pickup date, return date, price,
  car type, and supplier rating.
- View detailed car information, including vehicle specifications, pricing,
  availability, images, and supplier rating.
- Verify customer identity using government ID, driving license, document
  validation, and face recognition.
- Book cars through a secure workflow that includes verification checks and
  payment processing.
- Let suppliers manage cars by adding, updating, removing, and verifying
  vehicle information and documents.
- Let suppliers approve, reject, track, and complete booking requests.
- Protect individual supplier privacy by hiding exact home addresses while
  still showing general location information.
- Support admin reporting for revenue, bookings, payments, users, and disputes.

## Domain Model

The documented domain model includes these main concepts:

- `Customer`: account details, profile updates, document uploads, and booking
  history.
- `Supplier`: supplier type, company information, rating, verification status,
  privacy settings, car management, booking approval, and reports.
- `Car`: brand, model, price per day, availability status, car type, images,
  details, and availability updates.
- `Booking`: pickup date, return date, total price, booking status, creation,
  cancellation, modification, confirmation, and duration calculation.
- `Payment`: amount, payment method, payment status, transaction ID, receipts,
  validation, and refunds.
- `Verification`: identity verification, document validation, face recognition,
  approval status, and verification timestamps.
- `VerificationDocument`: document type, document URL, upload date, expiry date,
  upload, validation, and deletion.
- `Review`: rating, comment, review date, add, edit, and delete operations.
- `Report`: report type, generated date, revenue, booking totals, summaries,
  exports, and revenue calculations.
- `Admin`: supplier verification, user removal, report generation, payment
  monitoring, and dispute resolution.

## Tech Stack

### Frontend

- Next.js
- React
- TypeScript
- Tailwind CSS
- Radix UI components

### Backend

- Go
- Gin HTTP framework
- GORM
- PostgreSQL
- Redis
- MinIO
- JWT authentication

## Project Structure

```text
.
├── backend/      # Go API, authentication, role routes, database, cache, storage
├── frontend/     # Next.js frontend application
└── Documents/    # Requirements, diagrams, scenarios, and acceptance criteria
```

## Documentation

Project requirements and design artifacts are stored in `Documents/`, including:

- `Problem Statement.docx`
- `Main Scenario.docx`
- `Acceptance Criteria.docx`
- `Use Case.pdf`
- `ClassDiagram.pdf`
- `sequence-diagrams.md`
- `User Story Map.jpg`
- `Use Case Diagram V2.png`

## Built By

| Student ID | First Name   | Last Name          |
| --------------------- | ------------ | ------------------ |
| 620423                | Amr          | Rizkallah          |
| 620419                | Nida         | Javed              |
| 620356                | Nguyen       | Thi Hoang Phuong   |
| 620283                | Darkhanbayar | Erdenebat          |
