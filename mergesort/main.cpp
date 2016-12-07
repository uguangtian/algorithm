/*
 * =====================================================================================
 *
 *       Filename:  main.cpp
 *
 *    Description:  mergeSort
 *
 *        Version:  1.0
 *        Created:  12/07/2016 12:04:52 AM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  YOUR NAME (),
 *   Organization:
 *
 * =====================================================================================
 */



void MSort(int SR[],int TR1[],int s, int t){
	int m;
	int TR2[MAXSIZE+1];
	if(s==t){
		TR1[s]=SR[s];

	}else{
		m=(s+t)/2;
		MSort(SR,TR2,s,m);
		MSort(SR,TR2,m+1,t);
		MSort(SR2,TR1,s,m,t);
	}
}
void MergeSort(SqList *L){
	MSort(L->r,L->r,1,L->Length);
}

