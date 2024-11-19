/* u s i n g System ;
namespace F i b o n a c c i R e c u r s i v e {
c l a s s Program {
s t a t i c v o i d Main ( s t r i n g [ ] a r g s ) {
Console . WriteLine ( Fib ( 5 ) ) ;
Console . WriteLine ( Fib ( 1 0 ) ) ;
}
s t a t i c i n t Fib ( i n t n ) {
i f ( n>0) {
i f ( n<=2) {
r e t u r n 1 ;
} e l s e {
r e t u r n Fib ( n=1) + Fib ( n =2);
}
} e l s e {
r e t u r n 0 ;
}
}
}
*/

using System;
namespace  FibonacciRecursive
{
    class Program{
        static void Main (string []args){
            Console.WriteLine(Fib(5));
            Console.WriteLine(Fib(10));
        }

        static int Fib(int n){
            if (n > 0){
                if (n <= 2){
                    return 1;
                } else {
                    return Fib (n = 1) + Fib(n = 2);
                }
            } else{
                return 0;
            }
        }
    }
}