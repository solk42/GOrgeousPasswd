# Worklog  
A little diary of my first steps with go.  

---  
## Day 1 - 2020-04-12
| time | message |
|-|-|
| 10:08:05 | Well, golang.<br>Let's see what's it all about on tour.golang.org | 
| 10:51:07 | Well...<br>No class statements... And Pointer stuff... getting nervous :-O | 
| 11:29:10 | First quick read of the basics done... <br>I need a break.<br>**Summary:** I have no idea what I'm doing | 
| 13:16:09 | Since I only have my Windows Tablet here, I'll dig out my dusty RasPi and set it up for golang... | 
| 14:03:46 | Headless RasPi is running a go "hello world".<br>VSCode with SSH plugin runs on the tablet.<br>First I have to take care of other things, maybe I'll continue tonight.<br>**Summary:** Basics done, ready for task-takeoff (later) | 
| 16:43:43 | Setup of Github repo done: https://github.com/solk42/GOrgeousPasswd<br>Added first task description as md file | 
| 16:49:00 | starting first attempt of basic generator function ... to be continued | 
| 17:03:39 | well... generating a not really random string works - need to lookup how to generate it more realrandomly :-/ | 
| 17:09:09 | that was easy - added rand.Seed with time function on each call.<br>**Summary:** First babystep of TASK-01-01 done, will continue later or tomorrow | 
| 17:49:30 | added a simple diary file from my (quick and dirty) coded PHP diary-note ;) | 

---  
 
## Day 2 - 2020-04-13
| time | message |
|-|-|
| 09:06:51 | Well. First I'll need some coffee. Then I'll have a look at slice handling in go<br>**Summary:** Planning to getting it to work, after that I'll need to cleanup (Hopefully I will have a better understanding of go then) | 
| 09:29:09 | :coffee: :coffee: :coffee:<br>brain is booting, do not turn off | 
| 10:13:38 | Basic solution for TASK-01-01 done.<br>Really need to cleanup my mess :-/<br>Later I'll need to get into go packages. | 
| 11:57:08 | One simple step done - moved helper function to separate package.<br>But general unsure if this is the right approach (folder and import structure) | 
| 12:29:21 | Moved configuration to generator package - I guess I'll need to refactor again..<br>**Summary:** Strategy: Learning by doing. Trying to clean up step by step | 
| 16:08:09 | Took a little break :bath: :game_die:<br>Some ideas of go get a little clearer. I'll try to cleanup the generator code until midday tomorrow to get TASK-01-01 done. | 
| 16:40:06 | Moved the generator functions to the generator package.<br>It gets a little cleaner.<br>Now I need to get into public/private/accessibility logic of go and add some inline documentation (*go doc*?) for at least exported functions and structs | 
| 18:08:22 | Well. Cleanup is done for now. At least the exported functions and structues have basic go doc information.<br>Maybe I'll have a look on how to add go doc information as md file.<br>At least some manual testing is required to define the first task as done. | 

---  
 
