.. raw:: latex

  \setcounter{page}{1}
  \renewcommand{\thepage}{\thepart-\arabic{page}}
  \renewcommand{\thechapter}{\thepart-\arabic{chapter}}
  \setcounter{chapter}{0}


.. _Introduction-v1:

Introduction
============

This specification defines the core code and services that are
required for an implementation of the Pre-EFI Initialization (PEI)
phase of the Platform Initialization (PI) specifications
(hereafter referred to as the "PI Architecture"). This PEI core
interface specification (CIS) does the following:

* Describes the basic components of the PEI phase
* Provides code definitions for services and functions that are architecturally required by the UEFI PI working group (PIWG)
* Describes the machine preparation that is required for subsequent phases of firmware execution
* Discusses state variables that describe the system restart type

See :numref:`organization-of-the-pei-cis` for more information.


.. _principle-of-inclusive-terminology:

Principle of Inclusive Terminology
----------------------------------

The UEFI Forum follows a Principle of Inclusive Terminology in building and maintaining content for specifications. This means that efforts are made to ensure that all wording is perceived or likely to be perceived as welcoming by everyone regardless of personal characteristics. In some cases, the Forum acknowledges that wording derived from earlier work, for example references to legacy specifications not controlled by the Forum, may not follow this principle. In order to preserve compatibility for code that reads on legacy specifications, particularly where that specification is no longer under maintenance or development, language in this specification may appear out of sync with the Principle. The Forum is resolved to work with other standards development bodies to eliminate such examples over time. In the meanwhile, by acknowledging and calling attention to this issue the hope is to promote discussion and action towards more complete use of Inclusive Language reflective of the diverse and innovative population of the technical
community that works on standards.


.. _organization-of-the-pei-cis:

Organization of the PEI CIS
---------------------------

This PEI core interface specification is organized as shown in
"Organization of the PEI CIS."
Because the PEI Foundation is just one component of a PI
Architecture-based firmware solution, there are a number of
additional specifications that are referred to throughout this
document.

.. list-table:: Organization of the PEI CIS
   :name: organization-of-the-pei-cis-table
   :widths: auto
   :header-rows: 1

   * - Section
     - Description
   * - Overview
     - Describes the major components of PEI including the PEI Services boot mode PEI Dispatcher and PEIMs
   * - PEI Services Table
     - Describes the data structure that maintains the PEI Services
   * - Services PEI Services
     - Details each of the functions that comprise the PEI Services
   * - PEI Foundation
     - Describes the PEI Foundation and its methods of operation and the PEI Dispatcher and its associated dependency expression grammar
   * - PEIMs
     - Describes the format and use of the Pre EFI Initialization Module PEIM
   * - Architectural PPIs
     - Contains PEIM to PEIM Interfaces PPIs that are used by the PEI Foundation
   * - Additional PPIs
     - Contains PPIs that can exist on a platform
   * - SPEI to DXE Handoff
     - Describes the state of the machine and memory when the PEI phase invokes the DXE phase
   * - Boot Paths Boot
     - Describes the restart modalities and behavior supported in the PEI phase
   * - PEI Physical Memory Usage
     - Describes the memory map and memory usage during the PEI phase
   * - Special Paths Unique to the Itanium Processor Family
     - Contains flow during PEI that is unique to the Itanium processor family
   * - Security SEC Phase Information
     - Contains an overview of the phase of execution that occurs prior to PEI
   * - Dependency Expression Grammar
     - Describes the BNF grammar for a tool that can convert a text file containing a dependency expression into a dependency section of a PEIM stored in a firmware volume
   * - TE Image
     - Describes the format of the TE executable
   * - TE Image Creation
     - Describes how TE executables are created from PE32 executables
   * - TE Image Loading
     - Describes how TE executables are loaded into memory


.. _conventions-used-in-this-document:

Conventions Used in this Document
---------------------------------

This document uses the typographic and illustrative conventions
described below.

.. _data-structure-descriptions:

Data Structure Descriptions
###########################

Supported processors are "little endian" machines. This
distinction means that the low-order byte of a multibyte data
item in memory is at the lowest address, while the high-order
byte is at the highest address. Some supported processors may
be configured for both "little endian" and "big endian"
operation. All implementations designed to conform to this
specification will use "little endian" operation.

In some memory layout descriptions, certain fields are marked
reserved. Software must initialize such fields to zero and
ignore them when read. On an update operation, software must
preserve any reserved field.

The data structures described in this document generally have
the following format:

**Structure Name:** The formal name of the data structure.

**Summary:** A brief description of the data structure.

**Prototype:** A "C-style" type declaration for the data structure.

**Parameters:** A brief description of each field in the data structure prototype.

**Description:** A description of the functionality provided by the data structure, including any limitations and caveats of which the caller should be aware.

**Related Definitions:** The type declarations and constants that are used only by this data structure.

.. _procedure-descriptions:

Procedure Descriptions
######################

The procedures described in this document generally have the
following format:

**ProcedureName():** The formal name of the procedure.

**Summary:** A brief description of the procedure.

**Prototype:** A "C-style" procedure header defining the calling
sequence.

**Parameters:** A brief description of each field in the
procedure prototype.

**Description:** A description of the functionality provided by
the interface, including any limitations and caveats of which
the caller should be aware.

**Related Definitions:** The type declarations and constants that
are used only by this procedure.

**Status Codes Returned:** A description of any codes returned by
the interface. The procedure is required to implement any
status codes listed in this table. Additional error codes may
be returned, but they will not be tested by standard compliance
tests, and any software that uses the procedure cannot depend
on any of the extended error codes that an implementation may
provide.

.. _instruction-descriptions:

Instruction Descriptions
########################

A dependency expression instruction description generally has
the following format:

**InstructionName** The formal name of the instruction.

**Syntax:** A brief description of the instruction.

**Description:** A description of the functionality provided by
the instruction accompanied by a table that details the
instruction encoding.

**Operation:** Details the operations performed on operands.

**Behaviors and Restrictions:** An item-by-item description of the behavior of each operand involved in the instruction and any restrictions that apply to the operands or the instruction.

.. _ppi-descriptions:

PPI Descriptions
################

A PEIM-to-PEIM Interface (PPI) description generally has the
following format:

**PPI Name:** The formal name of the PPI.

**Summary:** A brief description of the PPI.

**GUID:** The 128-bit Globally Unique Identifier (GUID) for the
PPI.

**Protocol Interface Structure:** A "C-style" procedure template defining the PPI calling
structure.

**Parameters:** A brief description of each field in the PPI
structure.

**Description:** A description of the functionality provided by
the interface, including any limitations and caveats of which
the caller should be aware.

**Related Definitions:** The type declarations and constants that
are used only by this interface.

**Status Codes Returned:** A description of any codes returned by
the interface. The PPI is required to implement any status
codes listed in this table. Additional error codes may be
returned, but they will not be tested by standard compliance
tests, and any software that uses the procedure cannot depend
on any of the extended error codes that an implementation may
provide.

.. _pseudo-code-conventions:

Pseudo-Code Conventions
#######################

Pseudo code is presented to describe algorithms in a more
concise form. None of the algorithms in this document are
intended to be compiled directly. The code is presented at a
level corresponding to the surrounding text.

In describing variables, a list is an unordered collection of
homogeneous objects. A queue is an ordered list of homogeneous
objects. Unless otherwise noted, the ordering is assumed to be
First In First Out (FIFO).

Pseudo code is presented in a C-like format, using C
conventions where appropriate. The coding style, particularly
the indentation style, is used for readability and does not
necessarily comply with an implementation of the *Unified
Extensible Firmware Interface Specification* (UEFI 2.0
specification).


.. _requirements:

Requirements
------------

This document is an architectural specification that is part of
the Platform Initialization Architecture (PI Architecture) family
of specifications defined and published by the Unified EFI Forum.
The primary intent of the PI Architecture is to present an
interoperability surface for firmware components that may
originate from different providers. As such, the burden to conform
to this specification falls both on the producer and the consumer
of facilities described as part of the specification.

In general, it is incumbent on the producer implementation to
ensure that any facility that a conforming consumer firmware
component might attempt to use is present in the implementation.
Equally, it is incumbent on a developer of a firmware component to
ensure that its implementation relies only on facilities that are
defined as part of the PI Architecture. Maximum interoperability
is assured when collections of conforming components are designed
to use only the required facilities defined in the PI Architecture
family of specifications.

As this document is an architectural specification, care has been
taken to specify architecture in ways that allow maximum
flexibility in implementation for both producer and consumer.
However, there are certain requirements on which elements of this
specification must be implemented to ensure a consistent and
predictable environment for the operation of code designed to work
with the architectural interfaces described here.

For the purposes of describing these requirements, the
specification includes facilities that are required, such as
interfaces and data structures, as well as facilities that are
marked as optional.

In general, for an implementation to be conformant with this
specification, the implementation must include functional elements
that match in all respects the complete description of the
required facility descriptions presented as part of the
specification. Any part of the specification that is not
explicitly marked as "optional" is considered a required facility.

Where parts of the specification are marked as "optional," an
implementation may choose to provide matching elements or leave
them out. If an element is provided by an implementation for a
facility, then it must match in all respects the corresponding
complete description.

In practical terms, this means that for any facility covered in
the specification, any instance of an implementation may only
claim to conform if it follows the normative descriptions
completely and exactly. This does not preclude an implementation
that provides additional functionality, over and above that
described in the specification. Furthermore, it does not preclude
an implementation from leaving out facilities that are marked as
optional in the specification.

By corollary, modular components of firmware designed to function
within an implementation that conforms to the PI Architecture are
conformant only if they depend only on facilities described in
this and related PI Architecture specifications. In other words,
any modular component that is free of any external dependency that
falls outside of the scope of the PI Architecture specifications
is conformant. A modular component is not conformant if it relies
for correct and complete operation upon a reference to an
interface or data structure that is neither part of its own image
nor described in any PI Architecture specifications.

It is possible to make a partial implementation of the
specification where some of the required facilities are not
present. Such an implementation is non-conforming, and other
firmware components that are themselves conforming might not
function correctly with it. Correct operation of non-conforming
implementations is explicitly out of scope for the PI Architecture
and this specification.

.. _conventions-used-in-this-document-1:

Conventions used in this document
---------------------------------


.. _number-formats:

Number formats
##############

A binary number is represented in this standard by any sequence
of digits consisting of only the Western-Arabic numerals 0 and
1 immediately followed by a lower-case b (e.g., 0101b).
Underscores or spaces may be included between characters in
binary number representations to increase readability or
delineate field boundaries (e.g., 0 0101 1010b or
0_0101_1010b).

A hexadecimal number is represented in this standard by 0x
preceding any sequence of digits consisting of only the
Western-Arabic numerals 0 through 9 and/or the upper-case
English letters A through F (e.g., 0xFA23). Underscores or
spaces may be included between characters in hexadecimal number
representations to increase readability or delineate field
boundaries (e.g., 0xB FD8C FA23 or 0xB_FD8C_FA23).

A decimal number is represented in this standard by any
sequence of digits consisting of only the Arabic numerals 0
through 9 not immediately followed by a lower-case b or
lower-case h (e.g., 25).

This standard uses the following conventions for representing
decimal numbers:

* the decimal separator (i.e., separating the integer and fractional portions of the number) is a period;
* the thousands separator (i.e., separating groups of three digits in a portion of the number) is a comma;
* the thousands separator is used in the integer portion and is not used in the fraction portion of a number.

.. _binary-prefixes:

Binary prefixes
###############

This standard uses the prefixes defined in the International
System of Units (SI) (see
http://www.bipm.org/en/si/si_brochure/chapter3/prefixes.html )
for values that are powers of ten.

.. _si-prefixes:


.. table:: SI prefixes
   :widths: auto

   ============= ============= ==== ======
   Factor        Factor        Name Symbol
   :math:`10^3`  1,000         kilo K
   :math:`10^6`  1,000,000     mega M
   :math:`10^9`  1,000,000,000 giga G
   ============= ============= ==== ======


This standard uses the binary prefixes defined in ISO/IEC
80000-13 *Quantities and units -- Part 13: Information
science and technology* and IEEE 1514 *Standard for Prefixes
for Binary Multiples* for values that are powers of two.


.. _binary-prefixes-1:


.. table:: Binary prefixes
   :widths: auto

   ================= ============= ==== ======
   Factor            Factor        Name Symbol
   :math:`2^{10}`    1,024         kibi Ki
   :math:`2^{20}`    1,048,576     mebi Mi
   :math:`2^{30}`    1,073,741,824 gibi Gi
   ================= ============= ==== ======


For example, 4 KB means 4,000 bytes and 4 KiB means 4,096
bytes.


