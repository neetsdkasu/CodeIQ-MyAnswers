functor import System(show:S)Application(exit:X)define proc{F N A B C D E R}if N<100then{S R}{F N+1 D+E A+E B C D R+A}end end in{F 0 0 1 0 0 0 1}{X 0}end