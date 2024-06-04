package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestPredict(t *testing.T) {
	// create a ResponseWriter for testing
	w := httptest.NewRecorder()
	prompt := "based on this document, list the number of people who went on this trip."
	str, err := predict("llama3", prompt, testdoc, w)
	if err != nil {
		t.Errorf("predict() error = %v", err)
	}
	fmt.Println("str:", str)
}

const testdoc = `US TRIP AGENDA RSA Conference, Tech Meetings with CSPs and SGTS Partners
Salt Lake City, San Francisco & Seattle 30 April - 10 May 2024
_______________________________________
Delegation Government Technology Agency (GovTech) 1. Goh Wei Boon, Chief Executive of GovTech 2. Chang Sau Sheong, Deputy Chief Executive of GovTech
GDS 1. Kevin Ng, Senior Director, Gov Digital Services 2. Richard Tay, Senior Director, Government Infrastructure Group 3. Chia Hsiao Ming, Director, Gov Digital Services 4. Tang Bing Wan, Director, Gov Digital Services 5. Hunter Nield, Distinguished Engineer 6. Khi Ann Chua, Distinguished Engineer - Seattle only 7. Hudson Lee, Principal DevOps Engineer, Gov Digital Services - Seattle only 8. Liyana Muhammad Fauzi, Lead Product Manager, Gov Digital Services
CSG 1. Ong Hong Joo, Senior Director, Cyber Security Group - San Francisco only 2. Bernard Tan, Director, Cyber Security Group - San Francisco only 3. Kenneth Yap, Assistant Director, Cyber Security Group - San Francisco only 4. Gareth Yeo, Lead Cybersecurity Engnr, Cyber Security Group - San Francisco only 5. James Chua, Assistant Director, Cyber Security Group - San Francisco only
GIG 1. Richard Tay, Senior Director, Government Infrastructure Group 2. Arthur Lee, Director, Government Infrastructure Group 3. Lee Hoo Wah, Director, Government Infrastructure Group - San Francisco only 4. Boo Yong Nee, Deputy Director, Government Infrastructure Group 5. Gerald Lee, Deputy Director, Government Infrastructure Group - San Francisco only 6. Ian Ong, Assistant Director, Government Infrastructure Group - Seattle only 7. Joke Jong, Senior Project Manager, Government Infrastructure Group
Smart Nation and Digital Government Office (SNDGO) 8. Lua Rui Ping, Director - San Francisco only
Inland Revenue Authority of Singapore 9. Yap Mia Kai, CIO - Seattle, Microsoft EBC only
----------------------------------------------------------------------------------
Annex A - Routing Information Annex B - Accommodation Annex C - Talent Engagement
1


Meeting Overview

Mon

Tue

Wed

Thurs

Fri

Sat Sun

Mon

Tue

Wed

Thurs

Fri

SLC/SF

SF

SF/SEA

SEA

29 Apr

30 Apr

1 May

2 May

3 May

6 May (RSA) 7 May (RSA) 8 May

9 May

10 May

SIN to Salt SIN to SF CE, DCE & CE, DCE & CE, DCE &

CE & GIG CE & GIG CE, DCE & All

GIG

Lake City (Kev and HM)

(CE, DCE and GDS) Kev & HM

GDS Andreesse n Horowitz: 9am -

GDS GitLab Mtg A: 10am to

GDS Google (9am to 5pm) (bf &

1. Meeting with Pentera
2. Meeting

1. Meeting GDS with AWS

CrowdStr9am to 5pm

ike

(Lunch

MSFT EBC Day-1 (Entire day) 9am to 5pm

MSFT EBC Day-2 (Entire day) 9am to 5pm

DX -

12.45pm (bf 11am

GetDX.com & lunch

lunch available)

with MSFT 2. Meeting available)

Security

with Palo

DCE & GDS

available) Mtg B:

3. Meeting

Alto 6pm: Dinner

Cloudflare

1130am 1250pm

6pm: Dinner with

with Tenable

3. Meeting with Talent with

SEA to SIN: DCE, Kev &

1:15pm to (lunch avail) Talent

(CE, Arthur,

Trellix

HM

4:00pm

CSG, GIG @ (CE, Arthur,

Augment

Elastic 2pm to 5pm

SF)

CSG, GIG @

SF)

*the rest are

(optional)

DCE &GDS

leaving on Sat

5.00pm to 6.00pm

Notion or Analyse

PWC Global SFO to SEA:

Leaders

CE & GIG

(tbc)

10am (tbc) to

5.30pm to

1pm (bf

DCE &GDS @

6.30pm

available)

SEA MSFT:

Codeium: 9am to 5pm

1.30pm to

3.30pm

SFO to SEA: DCE & GDS

2


============================================= Tues 30 Apr - GDS
9:00am - 12:00pm: Meeting with DX @ Salt Lake City
Attendees from DX 1. Abi Noda, CEO 2. Greyson Junggren, CRO

Time
9:00am 9:30am

Topic
Introductions Government Technology Agency of Singapore Overview Sharing on EPP and Interests in Developer Productivity

Kevin

9.30am 10.00am
10.30am 11.30am 12.00pm

GetDX Platform Demo
GovTech POC and Integrations Discussion -in addition to the survey, to explore possibility of using DX to capture customenr satisftion within an SGTS product
Discussion on DX' research into developer productivity & how GovTech can be a strategic partner
Sharing on Stack and Possible Participation by DX
End and Next Steps

Gresysson Greysson Abi Noda Hsiao Ming

Key objectives 1. Engage GetDX expertise to show case value of EPP to senior stakeholders and agencies, and to assess
the priority areas to work on for our development teams to increase productivity 2. Discuss potential for customization of GetDX's solutions to specifically suit the GovTech's needs and
explore collaborative projects to innovate on improving developer experiences. 3. Evaluate GetDX's pricing structures to find a cost-effective solution that aligns with the GovTech's budget
constraints and long-term financial planning for technology investments.
Company Background GetDX has developed a comprehensive platform focused on enhancing developer productivity by integrating qualitative and quantitative insights. They aim to move beyond traditional metrics like lead times or pull requests to offer a more nuanced view of developer experiences. Their solutions include DevEx 360 for measuring internal developer experience, Data Cloud for engineering data analytics, and PlatformX which provides real-time intelligence for platform teams.
Recently, GetDX introduced several innovations and updates. They have launched a feature that allows for the correlation of qualitative and quantitative metrics, enhancing the depth of insights available to organizations. They also announced the general availability of PlatformX, a platform designed to give teams better real-time intelligence.
Prior Interaction - POC with DX for GDS and PIC:

3


OGP had conducted a POC with DX and had found it useful. Based on the good feedback from OGP, GDS is also embarking on a Proof of Concept (POC) that aims to implement the DevEx 360 tool, specifically the DX DevEX Cloud, for a one-month trial involving 100 developers across GDS and PIC. The POC is being conducted to assess the effectiveness of DevEx 360 in improving Developer Productivity within the organization. By leveraging the tool's features, such as research-based metrics, industry benchmarks, automated recommendations, and integration with collaboration platforms, the POC seeks to gather qualitative insights into teamwork and provide suggestions for improving team performance. Ultimately, the goal is to validate whether the interventions are effective in enhancing Developer Productivity and to identify any gaps that may be hindering development teams from performing at their full potential. : The POC is at no cost, though the annual cost per user DX DevEX Cloud is at USD408 per developer. For our POC, we are targeting 100 Developers across GDS for just one month, hence the estimated value would be USD3,400.
Potential Talking Points: 1. Seek Abi's expertise on showcasing value of EPP to senior stakeholders and agencies and how to
leverage DX to understand team's priorities for improvements. 2. Enquire about the possibility of integrating DX with our products like SHIP-HATS' GitLab, APEX, Elastic
etc. In order to include a more regular customer satisfaction feature that is trigger upon the completion of an action 3. Express the intention to establish a long-term partnership with DX, emphasizing the initial use of their platform for Proof of Concept (POC) to explore its potential benefits and compatibility with our needs. 4. Address the concern regarding the platform's cost by proposing a discussion on pricing flexibility, highlighting the scale of potential users (8,000 developers) which could justify a more favorable pricing structure. 5. Express interest in contributing to DX's research on improving developer experiences, suggesting a collaboration that could lead to mutual benefits and advancements in the field. 6. Extend an invitation to Abi Noda to participate in the upcoming Stack event in November, underscoring the value of his expertise to the event's success and audience engagement.
============================================
Wed 1 May - CE + DCE + GDS 9:00am - 12.15pm: Meeting with Andreessen Horowitz San Francisco -- 180 Townsend Street, San Francisco CA 94107
Refer here for Company Details.
7A.tte[nAdlseoesmet in 2022] 8. 9A.ndreessen Horowitz 1. Jesse Holmes (Partner, Go to Market Network) 2. Leila Hay (Partner)
Socket 1. Feross Aboukhadijeh (CEO)
Pinecone 1. Brian Bailey (Director of Sales, Enterprise) 2. Kyle Himmelwright (VP of RevOps)
Sourcegraph 1. Mark McCauley (Field Chief Technology Officer)
4


Reforge 1. Fareed Mosavat (Chief Development Officer) 2. Kevin Weiss (Head of Reforge for Teams)
Figma Patrick Delaney (Manager, Product Marketing)

Time
8.45am
9:00am 9:10am
9:10am 9:20am
9.25am - 9:55am
9:55am - 10.05am

Topic Light Breakfast available A16Z Introductions Government Technology Agency of Singapore Overview
a16z Introduction
Cybersecurity 9:25am - 9:55am Socket
Break

10.05am - 11.05am

AI 10:05am - 10:35am Pinecone 10:35am - 11:05am Sourcegraph*

11.05 am - 11.15am

Break

11.15am - 11.45am

Education & Community - Reforge

11:40am 12:00pm
12:00pm 12:30pm

Lunch & Wrap-Up Collaboration - Figma

5

Leader
Kevin
Feross Aboukhadijeh (CEO)
Pinecone 1. Brian Bailey (Director of Sales, Enterprise) 2. Kyle Himmelwright (VP of RevOps) Sourcegraph Mark McCauley (Field Chief Technology Officer)
1. Fareed Mosavat (Chief Development Officer) 2. Kevin Weiss (Head of Reforge for Teams)
Patrick Delaney - Manager, Product Marketing


12:30pm 12:45pm

Wrap Up & Next Steps

Key objectives 1. Gain insights into the innovative technologies and solutions being developed by the portfolio companies.
This involves diving deep into product offerings, its technology stack, and understanding how it fits into the broader ecosystem. 2. Explore potential partnership opportunities with Sourcegraph and other companies within Andreesen Horowitz's portfolio. 3. Learn about market trends, industry challenges, and emerging opportunities within the tech sector. By engaging with Andreesen Horowitz's companies, investors, and experts, visitors can gain valuable insights that can inform their own strategies and decision-making processes.

Company Background
Socket: Socket has recently raised $20 million to enhance the security of open source software, emphasizing its commitment to addressing vulnerabilities and risks associated with the use of open source technology in corporate environments. This funding aims to expand their team and support for various programming languages and integrations, reflecting Socket's response to the growing dependency and security concerns in open source software within tech industries
Pinecone: Pinecone is a vector database company that provides advanced solutions for managing and utilizing data in AI applications. Recently, Pinecone raised $100 million in a Series B funding round to support its growth and further develop its AI infrastructure technologies. The company has been recognized for its innovative contributions to the generative AI market, enabling powerful data management solutions. Pinecone has also announced a strategic partnership with Cloudera to accelerate the development of AI-powered applications, enhancing the way organizations utilize AI to improve operational efficiency and customer experiences.
Sourcegraph: Sourcegraph specializes in universal code search technology, enhancing developer productivity by making it easier to navigate and understand complex codebases. It offers features that enable code search across multiple repositories, enhancing code review processes and aiding in bug tracking and software maintenance.
Reforge: Reforge offers professional development programs and learning resources specifically tailored for professionals in the tech industry. The company focuses on high-impact areas such as growth, product, and engineering, providing in-depth insights and strategies that are directly applicable to real-world business challenges. Reforge collaborates with industry leaders to deliver cutting-edge content and networking opportunities to its members.
Figma: Figma, founded in 2012 by Dylan Field and Evan Wallace, emerged as a leader in web-based interface design tools emphasizing collaboration and user experience. Initially experimenting with different software ideas during their college years, they focused on creating a platform that would democratize design through accessibility and real-time collaboration capabilities. Notably, Figma introduced FigJam, a digital whiteboarding feature, and has consistently expanded its features to support educational and professional design needs. The company's innovative approach has positioned it prominently within the design community
Prior interaction: GovTech had visited Andressen Horowitz on previous trips but with a different set of companies.
6


Current and future contract value: The only company that we have commercial engagements with is Figma. GovTech has procured Figma at approx./ $1million across 5 years and are anticipating an increase over the next two years. DSAID, NDI, GDS, Workpal, PIC, and CSG are currently using Figma for designing wireframes and prototypes. For WOG, DWP is also considering opening up Figma to WOG.
For Reforge, there are some interests from the PM Community in GovTech. Teams like NDI were looking at purchasing Reforge license for about 10 pax - the cost would be $19,950 for a year, as individual membership is priced at US$1,995 per person annually. This membership allows access to two live courses per year and unlimited access to all their digital content and resources.
Potential Talking Points: Of the five companies, the three that we are most interested in are Sourcegraph, Reforge and Figma:
1. Ask Sourcegraph to demonstrate their advanced cross-project search functionalities and Cody (their code assistant) and explore how these compare with what we have in SHIP-HATS i.e. GitLab's Advanced Search function, GitLab Duo. Discuss how their tool can support our Engineering Productivity Program, particularly in facilitating code inner-sourcing and enhancing developer access across various agencies.
2. Request examples of Sourcegraph's successful implementations, particularly in environments of similar scale to ours. Solicit advice on fostering a culture of openness and collaboration, and how to effectively make reusable code accessible across governmental projects as part of innersourcing.
3. Discuss the availability of Reforge's enterprise licenses or custom courses that could be tailored to our needs. Explore scalable and cost-effective solutions for onboarding. So far, the NDI team has expressed interest in Reforge, and as they have found it to be useful, we could consider enabling access for additional product teams within GovTech.
4. Propose potential partnerships or special programs that Reforge could offer to enhance the capabilities of our product management or other core competencies. Emphasize the need for comprehensive training resources that could be deployed across GovTech to elevate our practices or incorporate as a Digital Academy offering
==========================================
Wed 1 May - CE + DCE + GDS 1.15pm - 4:00pm: Meeting with CloudFlare 101 Townsend St. San Francisco, California 94107, US
1A0tt.e[nAdlseoesmferot min C20lo2u2d] Flare 11. 12. 1. Alex Dyner, SVP, Special Projects
2. Nitin Rao, Chief Product Officer 3. Barry Fisher, Director, Product Marketing 4. Oliver Roup, Product VP 5. Dan Hollinger, Product Manager 6. Phillip Jones, Sr. Product Manager 7. Kenneth Lai, ASEAN Sales Leader 8. Lucas Tan, Account Manager, SG Public Sector
7


Time

Topic

1:00 pm Welcome & Introductions

Leader Mark (remotely)

1:15 pm

Cloudflare vision & Strategy Overview of the Connectivity Cloud story High-level trends we're seeing in the world 1-2 customer stories

Alex + Mark

1: 35 pm

Voice of Customer GovTech will share more about SEEDs plans on how they would like to segment security profiles to expand to other groups of users

GovTech

2:00 pm Product Overview & Roadmap Nitin Goal: understand the breadth of our solution

2:30 pm

Security Everywhere
 Adopting a Everywhere Security approach that delivers secured access to its workforce across web and multi cloud environments
 Deep dive into Zero trust solution and roadmap  Share how CF ZT caters to different market segments (i.e.
now GovTech is only using them for developers but are also looking for ideas which other groups would find this useful as well. Such as offshore users, Public officers etc.)  Threat intelligence

Barry

3:00 pm

Multi-tenancy and Failover roadmap
Ongoing monthly discussions with ZT PMs on these 2 topics. GovTech is requesting for an update and commitment from Nitin at the executive level - how and when these will be delivered

Oliver + Dan

3:20 pm

Developer Services High level of developer services. Particularly concerned about the security risk of having all storage on only one vendor (AWS S3), would like to explore how to disseminate the risk by having 2 vendors - CF R2.

Phillip Jones

4:00 pm Executive Closing

Nitin

New feature requests:  Multi-Tenanted Services
8


 Enhancement of operational resilience by continuing initiatives like Code Orange
Background Info Prior interaction: At working level, we have our weekly cadence meeting with Cloudflare on current issues, feature requests, financial issues as well as technical discussions.
On smaller scale feature requests, Cloudflare reps like Annabelle and Jeremy would help to relay the message back to HQ which is often tracked and discussed at the monthly PM meetings. For bigger items (e.g. Multitenancy and failover roadmap), they are more often than not discussed at the monthly executive meetings.
Current and future contract value: Based on our last invoices we are spending US$75,636.00/mth (approx. US$907,632.00.00 per year). We foresee the spendings with Cloudflare will grow beyond US1M per year, as we expand SEED to more groups of users, providing multi-tenanted tunnels for our tenants' tenant and eventually when we join hands with CStack to provide CDN services for their users. Other possible new use-cases that will push exponential growth with Cloudflare are related to GCC initiatives (e.g. Synapxe's HCC integration, onboarding of development environment from system integrators such as NCS).
Potential Talking points
1. Commit Cloudflare to prioritize on our multi-tenant self-service needs for agencies to onboard to cloudflare for access to both GovTech and agencies' resources. Share plans to integrate external IAM systems with SEED and explore using Cloudflare as an IdP broker to enhance security and user experience. Address VPN limitations by considering multi-tenant options and centrally managed configurations. Discuss the lack of trust in the current IdP (Entra ID) and plans for a full IdP from Cloudflare (as an replacement or resilient failover to Entra). Also, explore how Cloudflare deals with token theft and replay.
2. Discuss the expansion and segmentation of security profiles to cater to diverse user needs. This includes introducing segmented profiles such as SEED classic, SEED for selected government tools, SEED+ for high-risk users, and SEED lite for BYOD users. Discuss implementing Cloudflare's RBI to enhance security measures and manage tokens for restricted access. Evaluate using Microsoft's Continuous Access Evaluation to combat token theft and ensure secure authentication. Additionally, address the current/evolving landscape of scams and phishing seen by Cloudflare and the mitigations/tools available.
3. Discuss how Cloudflare could enhance operational resilience by continuing initiatives like Code Orange. Advocate for robust 24/7 TAM support as SEED expands its global reach.
4. Seek to streamline the procurement process with SEED, possibly through new approaches like public listings on CSP marketplaces or Enterprise Agreements, ensuring transparency, fairness, and value for money.
5. Explore next-generation serverless cloud platforms to reduce cost and accelerate delivery for workers and AI platforms.
9


===========================================
Thurs, 1 May - Optional Attendance from GovTech 5.00pm - 6.00pm, Meeting with Augment tbc
1A3tt.e[nAdlseoesmferot min A20u2g2m] net 11.4. xx 15.

Time 5pm 5.10pm 5.30pm 5.45pm 6.00pm

Topic
Intro to GovTech and EPP Sharing on Augment's Feature Demo Closing End

Leader Kevin

Key objectives
1. Explore the specifics of Augment's technology with a focus on its ability to adapt to contextual knowledge.
2. Evaluate how well Augment's coding assistant could integrate with the government context, especially with the need for data residency
3. Identify potential areas for strategic partnership or collaboration (e.g. a POC that could help us understand a learning model that best suits the government context, or help us study developer productivity further) that could mutually benefit both Augment and GovTech
Background Augment, a new player in the AI-powered coding platform market, recently launched with significant backing, including from former Google CEO Eric Schmidt. The company came out of stealth mode with $252 million in funding, valuing it near a unicorn status at $977 million post-money. Augment is designed to enhance the productivity and quality of software development using generative AI technologies. It aims to address common industry challenges such as fragile software, complex maintenance, and lengthy feature development backlogs.
Overview: Augment is an AI pair programmer tool, similar to Github Copilot, but with a fundamentally different approach and underlying technology. Augment just exited stealth with the largest round of funding by any of the tools in this space with a $227M series B at ~1B valuation.
About the tech

10


From a feature set perspective, Augment is at parity with Copilot and all tools in the market. Current tools like Copilot however, have approached the solution by outsourcing the AI stack to Open AI. While Open AI is a great general purpose AI model, it has no ability to understand or leverage context across a codebase. As a result, Copilot and other tools in the space are really only good for producing boiler plate code and will frequently hallucinate. They lack the ability to make references to underlying dependencies, API's, stylistic choices, classes, methods etc.

At Augment, we hired 25 of the best AI researchers in the world to purpose build the AI stack and entirely own the solution. By owning the solution, we were able to leverage a RAG (Retrieval Augmented Generation) based approach that makes Augment deeply codebase aware. That awareness produces significantly higher quality code and changes how AI tools can accelerate tasks like refactoring legacy code, testing, doing migrations, etc.

Some of the key advantages include:

 Providing much higher quality code through deep codebase awareness

 Much faster code suggestions (~3x faster than Copilot)

 Providing a much more secure solution

o We are SOC 2 compliant (Copilot is not)

o Contractually prohibited from training on any of your data or user actions

 Copilot like most solutions will use this approach in an attempt to gain context

of your codebase but it's far less effective than using Retrieval Augmented

Generation and poses a major threat to your code/IP

o

GDPR compliant

 Individual secure tenants for each customer

More

information:

https://docs.google.com/presentation/d/1EjbXDvULebd9SpZ9yjnAQyrbmujvhjx6HfBEu3f7W10/edit

#slide=id.g2b6108b3e74_0_0

Potential Talking Points
1. Since this is a start-up that is very keen on business, we have to be non-committal and engage in a more conversation, centring around learning more about the technology
2. To enquire if apart from investors, are there actual customers who are tryuing out Augment and have validated claims such as being "3x faster than Copilot"
3. To understand what are the success metrics for Augment that is different from other Coding Assistant tool providers and to suss out of their success metrics align with ours as well
4. To learn more about the technology and to understand how their LLM can support more contextualisation

11


============================================
Thurs, 2 May - CE + DCE + GDS 10:00am to 11:00am, 11:30am - 12:50am: Meeting with GitLab
2820 Scott St, San Francisco, CA 94123, US
1A6tt.e[nAdlseoesmferot min G20it2L2a]b 12.7. Sid Sijbrandij - Chief Executive Officer 13.8. Christopher Weber - Chief Revenue Officer 4. David DeSanto - Chief Product Officer (Virtual) 5. Johnathan Fullam - VP, Global Solution Architects 6. Sherrod Patching - VP, Customer Success Management 7. Fabian Zimmer - Director, Product Management 8. Craig Neilson - VP Sales APJ 9. Danny Petronio - Snr Major Account Executive APAC 10. Zane Chua - Snr Customer Success Manager

Meeting A - without Sid - on Commercials

Time

Topic

Leader

10.00am

Commercial Discussions on year 3 (2024-2025) renewal
Discrepancy over the licenses: Contracted for 2+1 years with 3770 licenses as 2nd year base. 3rd year pricing per contract to be locked in at USD641 per user per annum at the number set by GovTech that is above 4000 (e.g. 5500) - with only excess beyond that committed number to be charged at full-rate. GitLab is refuting this to have that price locked for up to 3770 - pending legal assessment. (delta of USD 0.5-0.7m)
10k architecture upgrade cost: Initial 10k architecture cost was not communicated before upgrade to address performance issues. Now, there is an additional 10k upgrade costs (approx. 900k USD) for year 3 onwards. This is a significant increase in the overall cost of GitLab (+20%). 5k architecture was not able to support when user count was around 3700. Would 10k be sustainable for 5k+ users? 5k architecture was mentioned by GitLab to support up to 5k users with headroom [link].
What are our planned commitments for year 3 onwards: FY24 - 5000 to 5500 FY25 - 6000 FY26 - 6500
Will we be able to keep to the ramp-up price?

Kevin, Christopher Weber - Chief Revenue Officer

Meeting B - Main Meeting

12


Time
11:20 am

Topic
Dedicated platform strategy  Feature enhancements timelines  Customer feedback on feature requests  Dedicated Runners features & timelines  Dedicated Runners pricing for budgetary purposes

Leader
Fabian Zimmer - Director, Product Management

11.40am

Gen Duo AI update & roadmap  Gen AI update & roadmap  Are there plans for data residency & when GitLab winning Google top DevOps Partner 2024

Johnathan Fullam - VP, Global Solution Architects,

12.00pm Discussion on Dedicated platform upgrade
12.30pm Lunch 12.50pm End

Christopher Weber - Chief Revenue Officer

Key objectives
1. Discuss support problems, resolve longstanding issues, and negotiate discounts and feature prioritization.
2. Enhance understanding of the organizational structure and strengthen relationships by sharing priorities and learnings.
3. Brainstorm partnership ideas and assess capabilities for resilience and AI integration.
Feature Requests:
 Introduce AI capabilities, such as Duo and extended MLOps features.
 Enable integration with JSM and ServiceNow for comprehensive DORA metrics capture.
 Develop a strategy for code scanning features to achieve parity with tools like SonarQube (for code quality) and FoD SAST and DAST.
 Expand the Enterprise Agile License to incorporate code review and approval functionalities.

Current Issues:  Pricing for GitLab hosted runners remains undisclosed.  Pricing Issues:
13


o Anticipated future pricing structures are concerning.
o Unexpected price increase to $10k architecture without prior notification.
 Beta runners lack essential functionalities and are not ready for general availability.
Delays in feature releases:
 GitLab advanced search was postponed from October 2023 to December 2023.
 GitLab pages release was delayed by three months, initially launched without user access control, making pages publicly accessible (from November 2023 to February 2023).
 Discussions on GitLab hosted runners started two years ago; however, the current beta version does not meet the needs of GovTech and its tenants due to missing critical features.
Background
GitLab Overview: GitLab, established in 2011, has evolved from a collaboration tool for programmers into a comprehensive DevSecOps platform that millions utilize to accelerate software delivery while enhancing security and compliance. The company has championed remote work and open-source practices, maintaining a strong commitment to improving the development lifecycle through frequent innovations.
Recent Developments: GitLab has recently launched Version 16.0, which includes new features such as Value Streams Dashboards, remote development workspaces, and enhancements to their AI-powered Code Suggestions. These additions are part of GitLab's ongoing effort to integrate more AI capabilities into their platform, aiming to boost efficiency across various stages of the software development process.
Current and future contract value: $3.4m with 4000 users at 5k architecture with intention to move towards 5-5.5k users in year 3. (GitLab wants to ramp this to $6.7m with 6500-7000 users at 10k architecture to retain $641 license price)
Potential Talking points with Sid
1. Engage Sid by introducing new GovTech leaders and informing him of Cheow Hoe's departure. Reestablish Kevin as the key leader looking after engineering productivity for Singapore government
2. Seek Sid's inputs on GovTech's intent to restructure into mission-focus and how Gitlab aligns teams working remotely on a joint mission.
3. (Sid has interest in Gov Tech space) Discuss challenges on how to be customer focus for our citizen and the ongoing challenge in dealing with scams
4. Thank Sid for Gitlab partnership and seek his commitment to partnership and to keep cost competitive.
Potential Talking points with GitLab team
1. Acknowledge and appreciate GitLab's ongoing support and collaboration, while addressing recent delays in feature rollout and requesting clarity on causes and preventative measures for future occurrences.
14


2. Emphasize ongoing market evaluation and the importance of GitLab recognizing consideration of competitive options, including GitHub.
3. Stress the importance of AI coding assistant in our space and for GitLab to determine if they would serve us or not.
4. Propose more favourable pricing and terms to facilitate the expansion of GitLab's solutions to additional government agencies, highlighting concerns about the increasing per user cost with the growth in user count.
5. Stress the intention to deepen the strategic partnership, highlighting the potential for expanding GitLab's adoption across more government agencies.
6. Discuss the need for GitLab's features to meet or exceed those of other market tools to ensure they remain the preferred choice, while also expressing interest in improving developer efficiency and deployment velocity, including strategies for innersource code sharing and the roadmap for remote developer environments to improve code security and reduce costs for VDI users such as IRAS, CPF, and Synapxe.
15


.

. =============================================

Thurs, 2 May - CE + DCE + GDS

2:00pm - 5:00pm: Meeting with Elastic

1A9tt.e[nAdlseoesmferot min E20la2s2t]ic

210. .Ash Kulkarni, CEO

21. .Santosh Krishnan, GM, Security

3. Gagan Singh, VP, Product Marketing, Observability

4. Lynn Shiow, Enterprise Account Executive

5. Jie-Hong Lim, Principal Solution Architect

Time (PST) 2:00 2:10 2:10 2:30 2:30 3:15
3:15 3:45
3:45 4:15
4:15 4:45
4:45 5:00

Topic

Agenda

Leader

Introductions

Welcome & Introductions

All

Voice of the Customer GovTech Priorities & Direction

Kevin Ng

Elastic's Key Highlights & Company Vision
Elastic Observability Roadmap and Enhancements Elastic Security Roadmap
Product Security
Q&A / Adjourn

 Our vision for the next 3 - 5 years  How do we deliver GenAI experiences?
focus*  What's Elastic Product organizational
archetypes and our approach to align our Product Management and Engineering across our platforms and solutions?
 Co-creation, Feedback & Enhancements  OpenTelemetry First vision & strategy  Enterprise-ready Serverless, Agentless
 Elastic Security for threat intelligence  Enterprise & Cloud Security  Modernised SecOps
How can enterprises adopt our product securely?
Wrap up

Ash Kulkarni, CEO
Gagan Singh
Santosh Krishnan
Mandy / InfoSec / ProdSec team All

Key objectives
1. While progress has been made, Elastic need to continue to strength support such as through dedicated support engineer offering, as issues are still numerous and takes time to resolve.

16


2. Outline GovTech direction to uplift agencies system resiliency and incident management practice, and explore multiple training/education channels to reach different audience types
3. Understand integration of GenAI experiences.
4. Clarify the organizational structure for product and engineering teams, address the rebranding of "Agentless" to "SaaS Agents," and enhance security measures for threat intelligence and cloud protection.
5. Focus on improving co-creation processes and elastic support, while strategizing secure adoption practices for enterprises.
Feature Requests: 1. Resolve the issue where an Identity Lifecycle Management (ILM) process fails when
a user is deleted, potentially a bug related to actions tied to the user's last interaction with the ILM. 2. Implement support for API key usage in the Watcher feature to enhance security for automated report generation and emailing. 3. Introduce the capability to create custom roles in Elastic Cloud accounts, beyond just Admin and Viewer, to include specific roles like Operator, Security, and Auditor. 4. Address the need for native agentless support in SaaS Agent to simplify deployment and management. 5. Enhance the Traffic Filter feature to support interconnectivity between different cloud providers using private networks. 6. Modify the Traffic Filter design to accommodate the usage of Real User Monitoring (RUM), which is currently incompatible.
Current Issues Faced: 1. Fragmented Documentation: The available documentation is disjointed and does not
provide comprehensive clarity. Specific issues include a lack of detailed information on permission settings and a shortfall in guidelines for complete end-to-end implementation processes, which are essential for effective system setup and use.
2. Restricted Support Access for Tenants: Tenants who utilize Kibana or Elastic Deployment are currently unable to raise support tickets directly. This limitation hinders their ability to resolve issues promptly and affects user satisfaction and system usability.
3. Limited Support Contacts: The mechanism for contacting support is confined to cloud account level controls, which restricts access to a select group of account members. This exclusivity can delay issue resolution and limits the availability of assistance when problems arise.
4. Challenges in Troubleshooting Elastic Integrations: There is a notable difficulty in obtaining effective support for troubleshooting Elastic integrations. This situation is often exacerbated by the complexity of the integrations and the specialized knowledge required to address such issues.
5. Insufficient Knowledge Among Support Staff: The support team shows a variable level of expertise concerning Elastic and Cloud technologies. This inconsistency in
17


knowledge can lead to ineffective problem-solving and guidance, prolonging downtime and frustration for users.
6. Unverified Support Recommendations: Often, the solutions and recommendations provided by the support team rae not pre-tested or validated. This practice can lead to the implementation of ineffective solutions, increasing the time and resources spent on issue resolution.
Background Info Prior interaction: CE had met Ash last October.
Current and future contract value The original contract value was at USD 945,000 over 3 years (inclusive of 10% discount) There were 3 top ups over the last 20 months, which total up to USD 2,597,400 (inclusive of original contract value). However, we noted that this value still unable to sustain the growth in StackOps adoption, we are also working to do a 4th top up by May/June 2024 to cater for the current growth and future which some large projects onboarding to StackOps - CPFB, GIG & HSS. We estimate the 4th top up to be USD 2,070,800 with additional options totalling up to USD3,129,600 (excluding the USD 2,070,800).
Potential Talking points:
1. Follow up on support issues raised since the last time CE met Elastic CEO. 2. Express gratitude for Elastic's partnership in supporting the Singapore Government in
its resiliency journey. Particularly, in addition to the platform, Elastic has been supportive of running workshops to enable WOG to learn more about monitoring and about the Elastic platform. They had also listened to our feedback on the issues with customer support, and arranged for the director-in-charge of customer success to discuss with us ways where Elastic can improve. 3. To establish strategic partnership in supporting Whole-of-Government (WOG) initiatives and emphasise the substantial increase in procurement value from $945k to over $3m for the next two years. This indicates a significant expansion in the scope and impact of the partnership, potentially leading to broader and more impactful collaborations within the government ecosystem. 4. Express that as our adoption increases, we look forward to exploring enhanced discount structures that align with the additional value. 5. To express the importance of a more integrated process to facilitate issue resolution, indicating a move towards a more streamlined and efficient support framework. Additionally, there is a transition from platinum to enterprise support, suggesting a shift towards more comprehensive and tailored support solutions. 6. To relay that the introduction of a new model for designated support engineers signifies a strategic shift in the support framework and should potentially lead to more personalized and specialized support for specific needs and challenges. 7. To request for more more detail and roadmap on donation of log schemas and profiling tools to the Opentelemetry foundation, and to discuss the viability of Opentelemetry as a common standard for security logs (in addition to observability)
18


19


============================================= Fri, 3 May - CE + DCE + GDS
9:00am - 4:00pm: Meeting with Google Google MP4 Building (Cloud Space)
1190 Bordeaux Dr, Sunnyvale, CA 94089, United States SVL-MP4-2-Code (30) [GVC]
2A2tt.e[nAdlseoesmferot min G20o2o2g]le 213. . Sachin Gupta Vice President and GM, Infrastructure 24. . Muninder Sambi - VP and GM, Networking and Security 3. Marco Genovese - Global Head of Chrome Enterprise Security Architecture 4. Rohan Grover, Director of Product Management 5. Andy Lee, Head of Public Sector 6. Hey Chee Weng, Customer Engineer

Infrastructure

Time 8:30 am 9:00 am
9:15 am 10:00 am 10:10 am
10:50 am

Topic
Registration and Welcome Refreshments
Introduction and Voice of the Customer
Empowering Al Cloud in the Box: Unleashing the Potential of GDCH
1. GDCH Updates 2. GDC-Engine-new 3. Partnership TPC Program
Break
The Age of Al: Goval From Dawn to Empowerment 1. Gen Al Gemini Update. 2. Future roadmap with LLM on GDCH (vs ChatGPT has made progress)
DORA Discussion Session [Remote] DORA Overview & Discussion

Leader
All Speakers Invited
Sachin Gupta - VP and GM Rohan Grover - Director of Product Management
Ting Liu - VP, Al Engineering
Nathen Harvey - Software Engineer

11:30 am

Lunch at Cloud Space Executive Dining (Would like some more options)

12:00 pm

Networking Discussion 1. Cross Cloud Interconnects 2. Cloud NGFW and NCC

Muninder Sambi - VP and GM, Networking and Security

20


12:40 pm

Chrome Enterprise GovTech has recently (Q1 2024) entered into an $18 million Enterprise Agreement with GCP. Additionally, we are working on a large Cyber Shield deal set for Q2.

Marco Genovese - Global Head of Chrome Enterprise Security Architecture

1:20 pm

Multi-cloud Security Posture Management systems

Jess Leroy - Senior Director, Product Management

2:00 pm 2:15 pm

Afternoon Break and Networking
GPU Updates GPU Roadmap/ Future Currently using 16 GPUs. Will expand to Global DC on GDCH: 100 GPUs $8-10M/year)

Rajeev Kalavar - Product Leader, AI Infrastructure

Option B

SRE Overview As GPU session is confirmed, shall we cancel SRE? Introduction to SRE. Exploring Google's strategies for SRE and system resiliency, critical for the continuous operation and reliability of Govtech systems.

GETCRE team will assign a speaker

3:30pm

Next Steps and Feedback

Account Team

Key objectives 1. Re-engage Google on GDCH as a GCC+ solution 2. Extend partnership by understanding how to leverage Google Gemini, particularly for code assistant 3. Understand the performance and cost characteristics for inter-cloud connectivity

Some issues we are facing:  Google Cloud Platform (GCP) adoption remained the lowest amongst the 3 main hyperscaler.  There are huge AI movements related to Google (programs such as AI Trailblazer is on its' second run),
however beyond testing/explorations the roadmap to possible production use-cases are still murky.  Local technical/engineering support is still lacking as compared to AWS and Microsoft. Case in point -
Google only feature one single Technical Account Manager (TAM) and one single Customer Engineer cum Solution Architect as opposed a much stronger (and larger) team of local support from AWS / Microsoft.  Local support (or lack of) is impacting billing operations and interventions were needed from GovTech.
New feature requests:
 Support for Gitlab integrations  Support for specific Kubernetes Metrics  Support for remote admin access (old issue, tracked since 2019)

21


Company profile and engagement with Singapore Government Google is building a third data centre in Singapore to meet the rapid user growth in the region, bringing its total long-term investment in such facilities here to US$850 million. The US$350 million facility is in Jurong West, and is close to its first two centres. In addition to powering consumer services, it also opened the 3rd zone Google Cloud Platform region in 2018 to serve enterprise customers and other clients.
Google Singapore 1,000 staff. On 22 Feb 2019, Google launched a developer space at its Singapore headquarters that will allow developers, entrepreneurs and community groups to network, train and grow their expertise. This space will be supporting initiatives such as housing community events and workshops, including a four-day Machines-Learning Boot camp. In Jan 2019, GovTech entered into an Enterprise Agreement (EA) with Google for a period of three years for Google Cloud services.
More recently in Mar 2024, GovTech entered into an early renegotiation of EA Service Schedule renewal with a much higher commit value of $18m over 3 years. Current cloudspend for GCP has increased significantly as compared to previous year, standing at around SG$5m+.
Prior interactions  April 2019: The meeting is a follow-up from a meeting with the MD, CTO Office, Will Grannis earlier in March 2019 where Cheow hoe and DS Kok Yam hosted the team at Hive. The meeting agreed that ongoing conversations between Google's OCTO and GovTech was important. The objective of the meeting was 1) to understand more about Google's latest Cloud Services Platform; 2) Zero trust strategy and; 3) use cases in the public sector to evaluate our use for it. Will also gave his commitment to speak at the STACK2020 next year and this would also serve as an ongoing cultivation effort.  August 2023: EBC Visit (CE Wei Boon and Cheow Hoe), hosted by Massimo Mascaro, Sr. Technical Director, (Office of the CTO) and Sachin Gupta.
Digital Technology Attachment Program (DTAP)  March 2019: Will and Cheow hoe discussed at the meeting in March about a possible exchange of engineers across GovTech and Google as part of the Digital Technology Attachment Program (DTAP). The DTAP is a programme initiated by GovTech to offer engineers in the Government the opportunity to gain industry exposure through a 6 to 12-month stint in global tech companies. So far, 4 positions from 2 companies (Pivotal and Cisco) are currently posted for staff application. There are ongoing talks between POG and Google (Angela).
Areas of Collaboration  In 2017, Google introduced the Cloud Machine Learning Engine to help developers with machine learning expertise easily build ML models that work on any type and size of data. It allows modern machine learning services, i.e., APIs--including Vision, Speech, NLP, Translation, and Dialogflow to be built upon pre-trained models to bring scale and speed to business applications. Kaggle, an online community of data scientists and machine learners, owned by Google LLC has grown to more than one million members and more than 10,000 businesses are using Cloud AI services,  The MOL team (Kok Hing), TMO and Hunter have engaged Google Singapore for deep dive sessions into the use of Dialogflow as we review the next steps for AskJamie.  GovTech has initiated the conversation on establishing a Collaboration Steering Committee with Google, co-chaired by Sew Bun to steer the relationship on collaboration on cloud but being new to the public sector, the response has been slower as compared to AWS & Microsoft.  In 2023, GovTech facilitated CSIT procurement of Google Distributed Cloud Hosted (GDCH) through GovTech's EA.
Google's projects with other Governments
22


 Google has been working with the US Government to use Google Cloud for many of their projects. Some of the use cases include the following:
 City of Los Angeles leveraging the cloud to design systems that visually share public safety and community services information during natural disasters to ensure citizens can easily find evacuation shelters, sandbag stations, local hardware stores, and real-time weather reports.
 Colorado moved to G Suite in Google Cloud in 2012 to simplify IT, transforming how teams work together and empowering employees to rethink inefficient business processes, realizing $32.5M in total cost avoidance as a result.
 National Institutes of Health leveraging on storage, compute, and machine learning technologies from Google Cloud to access biomedical data and accelerate research progress toward finding treatments and cures for the most devastating diseases.
 Mission Possible: Security and Critical Information Sharing in the Public Eye - A interactive panel session featuring major US City and State customers presenting an overview of the current security perceptions and challenges facing public sector organizations, highlighting forward-thinking leadership and flagship initiatives developed with Google Cloud to tackle large city and state-wide challenges.
Potential Talking points (TBC):
1. Appreciate Google's collaboration over the years and despite the low adoption, there was continuous push towards stronger partnership which was evident in the latest discussions with GovTech and SNDGO (helmed by Ruiping) on workstreams related to AI and Google Workspace (GWS).
2. We may want to stress the importance of continuous partnership to push for higher adoption, given the recent EA contract commitment of 18x higher ($18m) than the previous contract. The main vehicle for adoption is probably not towards typical services/systems migrated from on-premise systems (5 years into Cloud adoption journey across WOG, most Agencies have probably chosen their CSP platform of choice which is either AWS or Microsoft), but towards Gen AI or very specific security services.
3. Stronger efforts are needed to deepen strategic partnership beyond AI Trailblazer, given Google's headstart (via AI Government Cloud Cluster, or AGCC) in the AI Movement since last year, there should be more effort to focus joint GovTech/SNDGO- Google growth strategy to push for higher adoption and spend beyond POC and exploratory use-cases.
4. GovTech may want to push for more "investments" from Google in terms of local technical support which is severely lacking. Since 2019, they can muster no more than 3 to 4 engineers (max) to support us despite very strong projected growth. Even though we have onboarded a TAM since last year, we still have not instill stronger processes related to contractual/change/incident management.
5. There could be some strong push from Google for more GPU commitment (response could potentially be helmed by Ruiping).
6. There could be some potential push from Google on GDCH adoption across WOG (to be led by GovTech) beyond CSIT, on the GCC+ and confidential workload front.
7. There could be some potential asks from Google on the status of existing exploration and negotiation on GWS and Google Chronicle/Mandiant related services, with stronger push for more GWS adoption specifically. On Google Chronicle/Mandiant discussion, it could be deemed sensitive and CE may consider discussing this during RSA conference (led by Hong Joo).
8. Solution for remote administration still no progress since 2019.
23


============================================= Sun, 5 May - GIG + CSG
4:00pm - 9:00pm: RSA Microsoft Pre-event 2 New Montgomery St, San Francisco, CA 94105
2A5tt.e[nAdlseoesmferot min M20ic2r2o]soft 216. . Vasu Jakkal (CVP, Microsoft Security) 27. . Charlie Bell (Executive VP, Microsoft Security and Threat Protection)

Time

Topic

Attendees

Microsoft Pre-Day to kick-off

4pm

RSA Conference 2024 @

GIG, CSG

Palace Hotel

9pm

End

Agenda
Learn how your security teams can transcend individual expertise and protect more with AIpowered solutions.
Gain insights from Microsoft leaders, including Vasu Jakkal and Charlie Bell, and early access to the latest innovations in Microsoft Security through sessions, Q&A, and networking.
Key objectives: 1. Discover new generative AI-specific
controls can help you secure and govern the use of AI. 2. Learn how security and IT professionals can super charge their skills, collaborate more, see more, and respond faster with generative AI within Microsoft Copilot for Security. Network with peers.

24


============================================= Mon, 6 May - DCE + GDS
9:00am - 1:00pm: Meeting with PwC 405 Howard St Suite 600, San Francisco, CA 94105, USA 2A8tt.e[nAdlseoesmferot min P20w2C2] 219. . Christine del Rosario PwC Office Managing Partner for San Francisco 320. . Scott Likens Global AI Leader 3. Joe Harrington , PwC Innovation Hub 4. Bret Greenstein. PwC GenAI Leader (Remote) 5. Dane Ingram PwC Technology Team (Remote) 6. So Cher

Time 8.30am 9 am
9:15 am
10:00 am 10:45 am 11:15 am 11:45 am

Topic
Arrival and Breakfast
Welcome and Introductions
PwC's AI Journey  Our $1B investment  The PwC Factory Model  Our Responsible AI Foundation  Upskilling Our People  PwC's Client Zero Story  AI Survey / Market View  Market Outlook: Anticipated AI Impact by Industry
AI at PwC: Our Tech Roadmap  Rollout and Approach  Deployment and Adoption of AI Tools  ChatPwC Architecture

Leader
Christine del Rosario PwC Office Managing Partner for San Francisco Scott Likens Global AI Leader
Scott Likens Global AI Leader Ilana Golbin Blumenfeld PwC Responsible AI Lead
Scott Likens Global AI Leader Dane Ingram PwC Technology Team (Remote)

Break and Available Demos with PwC Staff (CoPilot & ChatPwC)
Use of AI in Coding & Development  Code Whisperer, GitHub and other Emerging Tools for Developers  Engineering Productivity & Output Implications  Potential job losses and disruption  Velocity and Pace of Change  Implications on Staffing Models
Open Discussion & Close

Scott Likens Global AI Leader Joe Harrington PwC Innovation Hub Bret Greenstein PwC GenAI Leader (Remote)
All

25
`
