/*
TODO:

Modulo
args[]
Statische variablen
!
no empty declaration
doubles im lexer? + "-" vor zahlen
keine klammern um if / else/ while mit einer anweisung

*/
using system;
using bba;
using asda;

namespace features{
    class features{

        /*static int i = 2;
        static bool sada = false;
        static string buu = ";";*/ 
        static void main (string[] args){
            // Comments - Special Characters?
            /* Multiline
            Comments+
            ! */

            // Literals
            int i = 1; // <- Integer
            string s = "a"; // <- String
            bool b = true; // <- Bool
            double d = -2.345; // <- Double

            // Operators:
            i = 1+2;
            i = 2-34;
            i = 23 * 2;
            d = 3/5;
            i = 27 % 4;

            // Comparisons
            b = (true == true);
            // <; >; <=; >=; !=

            // Bools
            b = true && false;

            // Operator Prädizidens



            // Variablen Deklaration nur am anfang des Codes
            // Strong Typesafety
            // No Explicit Casts
            
            // Assignments:
            i = 2;
            // Round down
            int j = 2.0;


            // If/ Else:

            if (true != false){
                int sum  = sum + 2;
            }else{
                // Optional!!
                int sum = sum +3;
            }


            // While
            int i = 0;
            while(i < 5){
                i = i +1;
            }

            // Geht auch ohne {}:
        
            while (i < 5)
            {i = a;}
            // kms

            // ## Methods



            // Method Calls:
            doFunc(1); // <- Verwirft Rückgabe Typ

            int i = doFunc();
            // Kein Polymorphismus!!

            // stdout:
            Console.WriteLine("hi");
            // Newline
            Console.WriteLine(doFunc(3));

        }
        
                    // Alle Methods sind static:
            // Return Types: int/ double/ bool/ string/ void
            // args sind optional
            static void doFunc(){
                // local variables nur am anfang definiert
                // Scoping!!
                bool b = true;
                if (b == true){
                    return; // Empty Returns work too
                }
                if (b == false){
                    return false;
                }
                // Void Function without return value
            }
    }
}