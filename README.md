# interview_challenges
Coding examples for technical interviews (Python &amp; Go) 

- Knowledge Engineer (Python): We want to see basic to intermediate programming skills (opening, creating a file, reading text and exporting data). We expect the candidate to be able to complete at least Level 1.

a) Level 1 - Basic: Files are open in memory. Possible issue, the file content cannot fit in memory

b) Level 2 - Progressive: Files a read line by line: Possible issues, the line is bigger than the memory

c) Level 3 - Class: Creating a class would demonstrate excellent organizational skills

 

- Senior Knowledge Graph Engineer (Golang): We want to see excellent understanding of concurrency, memory management and performance. (buffers, go routines, communication channels). It is expected the candidate to be able to complete at least level 3 (either variant is ok)

a) Level 1 - Basic: Files are loaded in memory and are read consecutively. The issue with that is reduced performance and possibility that the content would not fit in memory

b) Level 2 - Buffers: Creating buffers allows us to read the files efficiently without concerns of data fitting in memory

c) Level 3 - Concurrency: Files are open and processed concurrently. The use of input and output channels to synchronize the communication among go-routines help us improve the performance X times.

d) Level 3a – Concurrency wih Struct: Instead of synchronizing go-routines, a struct containing the word count map with a RWLock would allow us to create simpler to read code.