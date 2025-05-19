# Status update 

## Quick Intro to myself 

- I am Nabeel Khan, an IITG2024 graduate, and have been working at Joyride Games, for the past 10 months as a fullstack developer, and have worked across 2 teams.
- Initially Superchamps, wherein I was responsible for social quests, and as of late, I am part of the rather small team building a crypto decentralized exchange 
- Really liked what Farmako is building, in terms of the need of quick commerce in medicines. 
- I understand, from a user's pov why it is necessary to have a separate entity for quickcommerce around medicines/healthcare.
- From my personal experience as well, where in I tried getting stomach ache but failed to get them, cause the doctor couldnt verify if the medicine was apt for me or not.
- Plus other user pain points, like high availability of the dark stores, as compared to regular stuff, like bread, eggs or milk, medicines need to be available 24x7. 
- If not the right medicine, but hopefully something close enough, (yes, this is a joke, around CAP theorem being applicable in business...)
- I would really like to work on Farmako, and believe I can be a value contribution to the same.

- Based on my previous experiences, my strength lies in product thinking, and I believe I will be able to make a difference in Farmako by the same.

- Will be detailing out on what has been done and what all I have not done, and a possible explanation on how they can be implemented 
- So I was able to get the DB design done, through a discussion with GPT, I finalized to 7 entities, 

    - COUPONS 
    - MEDICINE_CATEGORIES
    - MEDICINES
    - COUPON_APPLICABLE_CATEGORIES
    - COUPON_APPLICABLE_MEDICINES
    - COUPON_USAGES 
    - TIME_WINDOWS
    - DISCOUNTS

- Made creation, validation (usage), and verification flow as well.
- Was definitely able to get the implementation done, with a certain few assumptions which are available in assumptions.md, although was not able to devtest things entirely
- Was able to write a docker file for building it into an image and running and testing things by hitting the remote port.
- Was able to add the swagger commands

- Dont mean to make an excuse off of thigns, but I am having a rather product release within a month, as a result I tried how much I could. Hope thats fine.

- Was not able to devtest edge cases completely
- Was not able to take care of goroutines or mutexes
- Was not able to implement request-scoped context, due to time constraints, and me focusing more on the core features.
- Was not able to implement caching patterns

- Especially in DB read-heavy or async operations, goroutines could significantly improve performance.
- Mutexes may become relevant in tracking coupon usage counts to avoid race conditions.
- Havent verified SQL queries, if there is DB level safety


