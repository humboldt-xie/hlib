/*
 * =====================================================================================
 *
 *       Filename:  c.cc
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  2018/06/05 23时56分42秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  YOUR NAME (), 
 *        Company:  
 *
 * =====================================================================================
 */
#include <stdio.h>
#include "c.h"
int global;

int int_func(int a) {
	return (a*100000)+(++global);
}

int main() {
    global=0;
    GoInt a = 12;
    GoInt b = 99;
    printf("awesome.Add(12,99) = %lld\n", Add(a, b));
    printf("awesome.Cosine(1) = %f\n", (float)(Cosine(1.0)));
    GoInt data[6] = {77, 12, 5, 99, 28, 23};
    GoSlice nums = {data, 6, 6};
    Sort(nums);
    for (int i = 0; i < 6; i++){
        printf("%lld,", ((GoInt *)nums.data)[i]);
    }
    GoString msg = {"Hello from C!", 13};
    Log(msg);

    CallBack(int_func);
}
