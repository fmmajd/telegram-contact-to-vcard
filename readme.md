# Convert Telegram Contacts to VCard

**TLDR: convert the html exported file of telegram contacts to an importable vcard file**

It was a foggy evening. All looked fine. I had no idea of the disaster ahead.

Calmly, I descended the cab, and went ahead my usual road.
Something felt wrong, but I shrugged the feeling and went on.

A podcast seemed appropriate, and I went in my purse to take out the phone.

O_O

It was not there.NOT. I looked and looked and looked. It was not.

Thanks to a miracle called "Find My Iphone", I found my iphone.
Shattered and bent and shredded, at the side of the autobahn.
Sitting there alone for about 10 minutes.
How many cars had run over my precious darling?
What could have been her last thoughts, doing everything she could to tell me where she was,
for the last time...


I was not much worried. Everything is synced, I told myself. Be rational.

Went home, began setting up a new phone, and then, I knew, ***OH THE PAIN***

Where are my contacts? Where are they?
I saw someone crazily running in the house, screaming in agony.
I put away my silly phone problems and went to help, only to see myself in the mirror.
Cliche, right?

Hope was not lost.
I logged in the icloud, and feeeeeeeeeew.
There they were, my dear contacts,

BUT

BUT

BUT

synced last a few years ago. Oooo

So I held my head up, and exported my contacts via Telegram.
I was happy.
Not the best option, but at least most of my treasure was restored, right?

Well, sort of

The format was not readable by my phone, so I sat, dug and wrote the script

**VOILA**

And here is my newly prepared v-card file to import in the new phone.
I enjoyed. You can too

## **How to?**

- clone the repository
  - ``git clone https://github.com/fmmajd/telegram-contact-to-vcard``

- install docker engine and docker compose
  - https://docs.docker.com/engine/install/
  - https://docs.docker.com/compose/install/

- rename the ``DataExport_<date>`` folder from telegram to ``data``  and put it inside the ``telegram-contact-to-vcard`` folder. So the structure should be: ``telegram-contact-to-vcard/data/lists/contacts.html``

- run docker container
  ``make up``

- run main.go file inside the container
  - first go into the container bash: `make bash`
  - then run: `go run main.go`
  - if you want to have a contact named 'Deleted Account' with all deleted account numbers of telegram, add a --add-deleted argument when running:
  `go run main.go --add-deleted`
  - Change the owner of the output file to your user
    - `USER=$(stat -c %g ./data/lists/)`
    - `chown $USER:$USER ./data/res.vcf`
    
  

- voila. your file is ready at data/res.cvf
